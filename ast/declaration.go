package ast

import (
	"bytes"
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Declaration struct {
	token *token.Token
	ids   []string
	ty    string
}

func NewDeclaration(token *token.Token, ids []string, ty string) *Declaration {
	return &Declaration{token, ids, ty}
}

func (declaration *Declaration) String() string {
	// 	declarations: declaration*;
	//  declaration: VAR ids type SEMICOLON;
	//  ids: ID idsPrime*;
	//  idsPrime: COMMA ID;
	//eg: var a, b, c int; var b bool; var f, g, h *myType

	var out bytes.Buffer
	out.WriteString("var ")

	for i, id := range declaration.ids {
		out.WriteString(id)
		if i == len(declaration.ids)-1 {
			out.WriteString(" ")
		} else {
			out.WriteString(", ")
		}
	}

	out.WriteString(declaration.ty)

	out.WriteString(";\n")
	return out.String()
}

func (decl *Declaration) BuildSymbolTable(errors []string, tables *st.SymbolTables, userTypes map[string]types.StructType) []string {
	// @memo declaration eg: var a, b, c int; var b bool; var f, g, h *myType

	// get declaration Type
	varType, varTypeName := types.StrToTy(decl.ty, userTypes)
	// @memo: StrToTy return UnknownTySig if type does not exist (i.e., not base type or not declared)

	// if type does not exist, append error
	if varType == types.UnknownTySig {
		errors = append(errors, SemErr(decl.token,
			fmt.Sprintf("unknown type `%s` in declaration.", varTypeName)))
	}

	// Iterate over all identifiers, check if already declared;
	// if so, throw error; else insert into `Globals` symbol table
	for _, id := range decl.ids {

		// if duplicate entry, append error and go to next
		if _, contains := tables.Globals.Contains(id); contains {
			errors = append(errors, SemErr(decl.token,
				fmt.Sprintf("duplicate variable declaration for `%s`.", id)))
			continue
		}

		// else, insert the variable in the symbol table
		// NOTE: this adds variables with `unknown` type to global var ST
		tables.Globals.Insert(id, st.NewVarEntry(id, varType, st.GLOBAL))
	}
	return errors
}

func (decl *Declaration) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	//eg: var f, g, h *myType -> ids = [f, g, h], ty = *myType

	// check if variable identifiers conflicts with function or struct identifiers
	for _, id := range decl.ids {
		if _, contains := tables.Funcs.Contains(id); contains {
			errors = append(errors, SemErr(decl.token,
				fmt.Sprintf("variable and function name conflict for `%s`.", id)))
		}
		if _, contains := tables.Structs.Contains(id); contains {
			errors = append(errors, SemErr(decl.token,
				fmt.Sprintf("variable and struct name conflict for `%s`.", id)))
		}
	}
	return errors
}

func (decl *Declaration) ToLLVM(tables *st.SymbolTables) {
	// create registers for global variables; fill symbol table
	for _, id := range decl.ids {
		// get var entry from symbol table
		varEntry, _ := tables.Globals.Contains(id)
		// create register
		varEntry.Reg = ir.NewGlobalRegister(varEntry.Ty, varEntry.Name)
		varEntry.Initval = SetInitVal(varEntry.Ty)
	}
}

func (decl *Declaration) PrintLLVM(tables *st.SymbolTables) string {
	var out bytes.Buffer

	for _, id := range decl.ids {
		varEntry, _ := tables.Globals.Contains(id)
		reg := varEntry.Reg.(ir.Operand)
		out.WriteString(fmt.Sprintf("%s = common global %s %s\n", reg.String(), reg.LLVMTy(), varEntry.Initval))
	}

	return out.String()
}



// Default values for int, bool, structs: 0, false=, null
// Obs: we are only using i64 types for now; obs: false = 0 in llvm.
func SetInitVal(ty types.Type) string {
	switch ty.(type) {
	case *types.IntTy:
		return "0"
	case *types.BoolTy:
		return "0"
	case *types.StrucType:
		return "null"
	default:
		panic("Invalid type to set initval for LLVM: " + ty.String())
	}
}
