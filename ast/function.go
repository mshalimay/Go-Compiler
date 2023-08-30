package ast

import (
	"bytes"
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strconv"
)

// TODO: change return type to types.Type

type Function struct {
	token        *token.Token
	id           string
	returnType   string
	parameters   *Parameters
	declarations []*Declaration
	statements   []Statement
	duplicate    bool
	cfg 		*ir.BasicBlock
	HasReturn 	bool
}

func NewFunction(token *token.Token, identifier string, returnType string, parameters *Parameters, declarations []*Declaration, statements []Statement) *Function {
	return &Function{token, identifier, returnType, parameters, declarations, statements, false, nil, false}
}

func (function *Function) String() string {
	// 'func' 'id' Parameters [ReturnType] '{' Declarations Statements '}'

	var out bytes.Buffer

	out.WriteString("func ")

	out.WriteString(function.id)

	out.WriteString(function.parameters.String())

	if function.returnType != "" {
		out.WriteString(" " + function.returnType + " ")
	}

	out.WriteString("{\n")

	for _, declaration := range function.declarations {
		out.WriteString(declaration.String())
	}

	for _, statement := range function.statements {
		out.WriteString(statement.String())
	}

	out.WriteString("}\n")
	return out.String()
}

func (f *Function) BuildSymbolTable(errors []string, tables *st.SymbolTables, userTypes map[string]types.StructType) []string {
	// create a funcEntry to populate the functions symbol table
	// @memo:
	// type FuncEntry struct {
	// 	Name 		string
	// 	RetTy 		types.Type
	// 	Parameters 	[]*VarEntry
	// 	Variables 	*SymbolTable[*VarEntry]
	// }

	// if duplicate function identifier, append errors and return
	if _, contains := tables.Funcs.Contains(f.id); contains {
		errors = append(errors, SemErr(f.token, "duplicate function declaration "+f.id))
		f.duplicate = true
		return errors
	}

	// if not duplicate, create a function entry
	funcEntry := st.NewFuncEntry(f.id, nil)

	// try to get return type and set it to funcEntry
	// @memo: StrToTy return UnknownTySig if type does not exist (i.e., not base type or not declared)
	retType, retTypeName := types.StrToTy(f.returnType, userTypes)

	if retType == types.UnknownTySig {
		errors = append(errors, SemErr(f.token,
			fmt.Sprintf("type `%s` of return of function `%s` does not exist.", retTypeName, f.id)))
	}
	funcEntry.RetTy = retType

	// fill funcEntry `Parameters` slice. @memo: `Parameters` is []*VarEntry
	addedParams := make(map[string]bool)
	for _, param := range f.parameters.params {
		id := param.id
		ty := param.ty
		// try to get the parameter type
		paramType, paramTypeName := types.StrToTy(ty, userTypes)

		// if type not found, append error
		if paramType == types.UnknownTySig {
			errors = append(errors, SemErr(param.GetToken(),
				fmt.Sprintf("type` %s` of parameter `%s` in declaration of `%s` does not exist.", paramTypeName, id, f.id)))
		}

		// if repeated parameter, append error and continue
		if _, ok := addedParams[id]; ok {
			errors = append(errors, SemErr(param.GetToken(),
				fmt.Sprintf("repeated parameter `%s` in declaration of `%s`.", id, f.id)))
			continue
		}

		// create a new VarEntry and add it to the funcEntry `Parameters` slice
		// NOTE: this will add parameters with `UnknownTySig` type to the funcEntry
		// NOTE: parameters scope is `LOCAL`
		funcEntry.Parameters = append(funcEntry.Parameters, st.NewVarEntry(id, paramType, st.LOCAL))
		addedParams[id] = true
	}

	// create a new symbol table for the function local variables, whose parent is the global symbol table
	funcST := st.NewSymbolTable[*st.VarEntry](tables.Globals)

	// iterate through `declarations` and add variables to the function symbol table
	// eg: if `var a, b int;` inside function, iterate over `a`, `b` and add to function symbol table

	// slice keeping track of added variables to check for duplicates
	addedVariables := make(map[string]bool)
	// NOTE: not using funcST for this bcs would look for variables in parent; alternatively, could change funcST/contains signatures

	for _, decl := range f.declarations {
		// try to get local variable type
		varType, varTypeName := types.StrToTy(decl.ty, userTypes)

		// if type does not exist, append error
		if varType == types.UnknownTySig {
			errors = append(errors, SemErr(decl.token,
				fmt.Sprintf("unrecognized type `%s` in local declaration of function `%s`.", varTypeName, f.id)))
		}

		// Iterate over variable identifiers and check if duplicate or = parameter identifier.
		// If so, throw error; else insert into symbol table.
		for _, id := range decl.ids {

			// if duplicate variable declartion, append error and go to next
			if _, contains := addedVariables[id]; contains {
				errors = append(errors, SemErr(decl.token,
					fmt.Sprintf("duplicate variable declaration `%s` in function `%s`.", id, f.id)))
				continue

				// if variable with same id as parameter, append error and go to next
			} else if _, contains := addedParams[id]; contains {
				errors = append(errors, SemErr(decl.token,
					fmt.Sprintf("variable `%s` in function declaration `%s` has same identifier as parameter.", id, f.id)))
				continue

				// if not duplicate, create a new VarEntry and add it to the symbol table
			} else {
				funcST.Insert(id, st.NewVarEntry(id, varType, st.LOCAL))
				addedVariables[id] = true
				// NOTE: this adds variables with unknown types to the local symbol table
			}
		}
	} // end of declarations loop

	funcEntry.Variables = funcST
	tables.Funcs.Insert(f.id, funcEntry)
	return errors
}

func (f *Function) TypeCheck(errors []string, tables *st.SymbolTables) []string {

	// check for conflicting identifier with global variables
	if _, contains := tables.Globals.Contains(f.id); contains {
		errors = append(errors, SemErr(f.token, "function and variable name conflict for `"+f.id+"`"))
	}

	// check for conflicting identifier with structs
	if _, contains := tables.Structs.Contains(f.id); contains {
		errors = append(errors, SemErr(f.token, "function and struct name conflict for `"+f.id+"`"))
	}

	// get the func entry to pass to the statements
	funcEntry, _ := tables.Funcs.Contains(f.id)

	// REVIEW: how to typecheck duplicate func declarations ?
	// they will look at the wrong func entry
	// for now, skip if duplicate using field `duplicate`
	if !f.duplicate {
		for _, stmt := range f.statements {
			errors = stmt.TypeCheck(errors, tables, funcEntry)
		}
		
		// check if function has return for all control paths
		var b bool
		warnings := make([]string, 0)
		hasReturn := false
		for _, stmt := range f.statements {
			warnings, b = stmt.HasReturn(warnings)
			hasReturn = hasReturn || b

			// After a return is reached, all following statements are unreachable
			if hasReturn {
				warnings = append(warnings, SemErr(stmt.GetToken(), "Unreachable code."))
			}
		}

		// If function doesn't has return for all control paths && is not void, error
		if !hasReturn && funcEntry.RetTy != types.VoidTySig {
			errors = append(errors, SemErr(f.token, "Function "+f.id+" missing return statement."))
		}
		f.HasReturn = hasReturn
	}
	return errors
}

func (f *Function) ToLLVM(tables *st.SymbolTables) *ir.BasicBlock {

	// get function entry in the symbol table
	funcEntry, _ := tables.Funcs.Contains(f.id)

	// fill symbol table with function's llvm name
	funcEntry.LLVMName = "@" + f.id

	//--------------------------------------------------------------------------
	// entry block
	//--------------------------------------------------------------------------
	entryBlock := ir.NewBasicBlock(ir.BASIC_ENTRY)
	
	// create register and allocate memory for return value; fill symbol table
	if funcEntry.RetTy != types.VoidTySig {
		// ret val register pattern: %_<f.id>retval
		retValReg := ir.NewVirtualRegister(funcEntry.RetTy, f.id + "_" + "retval")
		funcEntry.RetValReg = retValReg
		entryBlock.AddInstruction(ir.NewAlloca(retValReg))
	} else if funcEntry.Name == "main" {
		retValReg := ir.NewVirtualRegister(types.IntTySig, "main_retval")
		funcEntry.RetValReg = retValReg
		entryBlock.AddInstruction(ir.NewAlloca(retValReg))
	}


	// create register and allocate memory for parameters; fill symbol table
	for i, param := range funcEntry.Parameters {
		if i == 1 {
			entryBlock.AddInstruction(ir.NewComment("parameters"))
		}
		// param register pattern: %P_<paramname>
		paramRegister := ir.NewVirtualRegister(param.Ty, "P_" + param.Name)
		param.Reg = paramRegister
		entryBlock.AddInstruction(ir.NewAlloca(paramRegister))
	}

	// create register and allocate memory for variables; fill symbol table
	entryBlock.AddInstruction(ir.NewComment("local variables"))
	for _, varEntry := range funcEntry.Variables.GetEntries() {
		// local var register pattern: %<varname>
		varReg := ir.NewVirtualRegister(varEntry.Ty, varEntry.Name)
		varEntry.Reg = varReg
		entryBlock.AddInstruction(ir.NewAlloca(varReg))
	}


	// REVIEW this might be moved to loop upwards; doing here to preserve order in llvm instructions
	for _, param := range funcEntry.Parameters {
		// store parameter value in parameter register
		// store <ty> <value>, <ty>* <pointer>
		// eg: store i64 %initVal, i64* %P_initVal 
		entryBlock.AddInstruction(ir.NewStore(ir.NewVirtualRegister(param.Ty, param.Name), param.Reg.(ir.Operand)))
	}

	curBlock := entryBlock
	for _, stmt := range f.statements {
		curBlock = stmt.ToLLVM(tables, funcEntry, curBlock)
	}

	// REVIEW if f.HasReturn, might not need the final block. Consider removing it.
	// if main function, add instructions to the final block to return 0, if needed.
	if f.id == "main" {
		// store i64 0, i64* %_mainretval
		curBlock.AddInstruction(ir.NewStore(ir.NewLiteral(types.IntTySig, 0), funcEntry.RetValReg.(ir.Operand)))
		
		// %r41 = load i64, i64* %_mainretval
		reg := ir.NewTempRegister(types.IntTySig, funcEntry)
		curBlock.AddInstruction(ir.NewLoad(reg.GetLabel(), funcEntry.RetValReg.(ir.Operand)))

		// ret i64 %r41
		curBlock.AddInstruction(ir.NewRet(reg))
	
	} else if !f.HasReturn {
		curBlock.AddInstruction(ir.NewRet(nil))
	}


	f.cfg = entryBlock
	return entryBlock
}


func (f *Function) PrintLLVM(tables *st.SymbolTables) string {
	var out bytes.Buffer

	// get funcEntry
	funcEntry, _ := tables.Funcs.Contains(f.id)

	// define <ret_ty> @FUNC_NAME(<ty1> %PAR_NAME1, <ty2> %PAR_NAME2, ...)
	out.WriteString("define ")

	// return type
	if f.id == "main" {
		out.WriteString("i64")
	} else {
		out.WriteString(types.ToLLVM(funcEntry.RetTy))
	}
	out.WriteString(" ")
	
	// allocate memory for parameters
	out.WriteString(funcEntry.LLVMName)
	out.WriteString("(")
	for i, param := range funcEntry.Parameters {
		out.WriteString(types.ToLLVM(param.Ty))
		out.WriteString(" ")
		out.WriteString("%" + param.Name)
		if i != len(funcEntry.Parameters)-1 {
			out.WriteString(", ")
		}
	}

	// print all llvm instructions using visor's pattern
	out.WriteString(")\n{\n")
	out.WriteString(f.cfg.String())
	out.WriteString("}\n")
	return out.String()
}

func (f *Function) TranslateToAssembly(tables *st.SymbolTables) string {
	var out bytes.Buffer
	
	funcEntry, _ := tables.Funcs.Contains(f.id)

	out.WriteString("//" + f.id + "\n")

	//==========================================================================
	// Function header
	//==========================================================================
	out.WriteString(fmt.Sprintf("\t .type %s, %%function\n", f.id))
	out.WriteString(fmt.Sprintf("\t .global %s\n", f.id))
	out.WriteString("\t .p2align 2\n")
	
	//==========================================================================
	// Function prologue
	//==========================================================================
	
	// function label. Eg: `main:`
	out.WriteString(f.id + ":\n")


	// allocate bytes in the stack for all variables
	// 8 * local variables + 8 * parameters + 8 * temporary registers + fp + lr + retval
	nVar := len(funcEntry.Variables.GetEntries())
	totalSpace := 8 * nVar + 8 * len(funcEntry.TempRegisters) +
				  8 * len(funcEntry.Parameters) + 16 + 8

	// round to multiple of 16
	if totalSpace % 16 != 0 {
		totalSpace += 16 - (totalSpace % 16)
	}

	out.WriteString("// Prologue\n")
	out.WriteString(fmt.Sprintf("\t sub sp, sp, #%d\n", totalSpace))
	out.WriteString("\t stp x29, x30, [sp]\n")
	out.WriteString("\t mov x29, sp\n")

	//==========================================================================
	// Fill the symbol table with all llvm registers
	//==========================================================================
	// obs: naive stack based implementation
	
	// fill location for return register
	funcEntry.StackRegisters[f.id + "_" + "retval"] = "16"

	// fill location for local variables
	offset := 24
	for _, v := range funcEntry.Variables.GetEntries() {
		funcEntry.StackRegisters[v.Name] = strconv.Itoa(offset)
		offset+=8
	}

	// fill location for temporary registers
	for _, key := range funcEntry.TempRegisters {
		funcEntry.StackRegisters[key] = strconv.Itoa(offset)
		offset += 8
	}
	
	// fill location for parameter temporary registers
	for i, param := range funcEntry.Parameters {
		// for the temporary registers created in llvm with alloca
		funcEntry.StackRegisters["P_" + param.Name] = strconv.Itoa(offset)
		
		// for the parameters passed to the function
		if i < 8 {
			funcEntry.StackRegisters[param.Name] = fmt.Sprintf("x%d", i)
		} else {
			panic("Not implemented: More than 8 parameters; have to allocate to stack.")
		}
		offset+=8
	}

	// fmt.Println(funcEntry.StackRegisters)
	//==========================================================================
	// Function code
	//==========================================================================
	out.WriteString(f.cfg.TranslateToAssembly(tables, funcEntry))
	

	//==========================================================================
	// Epilogue
	//==========================================================================
	out.WriteString("// Epilogue\n")
	out.WriteString("\t ldp x29, x30, [sp]\n")
	out.WriteString(fmt.Sprintf("\t add sp, sp, #%d\n", totalSpace))
	out.WriteString("\t ret\n")
	out.WriteString("\t .size " + f.id + ", (.-" + f.id + ")\n")

	return out.String()

}

