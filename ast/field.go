package ast

import (
	"bytes"
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// Inherits from Expression
type Field struct {
	token 			*token.Token
	ids 			[]string
	ty 				types.Type
}

func NewField(token *token.Token, ids []string, ty types.Type) *Field{
	return &Field{token, ids, ty}
}

func (f *Field) String() string {
	// Field: a.b.c
	var out bytes.Buffer

	// out.WriteString("Field string method: ")
	for i, id := range f.ids {
		out.WriteString(id)
		if i != len(f.ids) - 1 {
			out.WriteString(".")
		}
	}
	return out.String()
}


func (f *Field) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{	
	// consider eg: a.b.c; var a *MyStru1; var b *MyStru2
	

	var varEntry *st.VarEntry

	// check if is parameter
	if entry, contains := funCallee.ContainsParam(f.ids[0]); contains {
		varEntry = entry
	} else if entry, contains := funCallee.Variables.Contains(f.ids[0]); contains {
		varEntry = entry
	} else {
		errors = append(errors, SemErr(f.token, fmt.Sprintf("%s, variable `%s` not declared", f.String(), f.ids[0])))
		f.ty = types.UndeclTySig
		return errors
	}

	// iterate over `a.b.c` and:
	// (i) check if `a` and `b` are pointers to structs
	// (ii) check if `MyStru1` and `MyStru2` are defined
	// (iii) check if `b` is field of `struct1` and `c` is field of `struct2`
	// return type of `c`

	// get `a` type
	currType := varEntry.Ty

	// iterate over fields `b` and `c`
	var fType types.Type = types.UndeclTySig

	for i, id := range f.ids[1:] {
		// iter 1: `a` is pointer to existing struct?; iter 2: `b` is pointer to struct?
		if !types.IsStructType(currType) {
			errors = append(errors, SemErr(f.token, 
				fmt.Sprintf("`%s`. Cannot access field of non-struct; `%s` is of type `%s`.", f.String(), f.ids[i], currType.String())))
			f.ty = types.UnknownTySig
			return errors
		}

		// iter 1: get `MyStru1` and its entry in the Structs symbol table
		strucName := currType.String()
		strucEntry, _ := tables.Structs.Contains(strucName)
		
		if strucEntry == nil {
			panic("Struct " + strucName + " not found in symbol table")
		}
		
		// 1st iter: `b` is field of `MyStru1`? 2nd iter: `c` is field of `struct2`?
		if _, ok := strucEntry.Fields[id]; !ok {
			errors = append(errors, SemErr(f.token, 
				fmt.Sprintf("%s. Field `%s` not found in struct `%s`", f.String(), id, strucName)))
			
			f.ty = types.UnknownTySig
			return errors
		}

		if i == len(f.ids) - 2 {
			// 3rd iteration: got to `c` get its type and break
			fType = strucEntry.Fields[id].Ty
			break
		}
		// update currType
		currType = strucEntry.Fields[id].Ty
	}
	// REVIEW loop above stops when find an error. Should continue looking for errors?

	f.ty = fType
	return errors
}
				

func (f *Field) GetType() types.Type {
	return f.ty
}

func (f *Field) GetToken() *token.Token {
	return f.token
}

// <result> = getelementptr <ty>, <ty>* <ptrval>, i32 0, i32 <index>

func (f *Field) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	reg := f.ToLLVMAux(tables, funCallee, block)
	newReg := ir.NewTempRegister(reg.Ty(), funCallee)
	block.AddInstruction(ir.NewLoad(newReg.GetLabel(), reg))
	return newReg	
}


func (f *Field) ToLLVMAux(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	
	// Consider below example:
	/* Type Point {
		a int;
		b string;
		c *Point2;
	}
	Type Point2 {
		x int;
		z *Point3;
		y int;
	}

	Type Point3{
		x int;
	}

	var p *Point;
	p.c.z.x = ...
	*/
	
	// get the var entry for `p`
	varEntry := GetVarEntry(tables, funCallee, f.ids[0])
	
	ty := varEntry.Ty				
	reg := varEntry.Reg.(ir.Operand)	
	for i:=0; i < len(f.ids) - 1; i++ {

		// get StructEntry for `Point` and `Point2`
		strucEntry, _ := tables.Structs.Contains(ty.String())

		// load the pointer to the struct into a new register
		reg1 := ir.NewTempRegister(ty, funCallee)
		block.AddInstruction(ir.NewLoad(reg1.GetLabel(), reg))

		//------------------------------------------------------------------
		// create GEP instruction to access the struct field
		//------------------------------------------------------------------
		// create a destination register
		reg2 := ir.NewTempRegister(strucEntry.Fields[f.ids[i+1]].Ty, funCallee)
		
		// Get the "position" for the accessed field among the struct fields
		// iter 0: Get "c" position among `Point` fields using: Fields["c"] =  {Ty: *Point2, Offset: 2}
		// iter 1: Get "z" position among `Point2` fields using: Fields["z"] = {Ty: *Point3, Offset: 1}
		// iter 2: Get "x" position among `Point3` fields using: Fields["x"] = {Ty: int, Offset: 0}
		fieldIndex := strucEntry.Fields[f.ids[i+1]].Offset

		// add gep instruction
		block.AddInstruction(ir.NewGep(reg2.GetLabel(), reg1, fieldIndex))

		// update reg and ty for next iteration
		// iter 0: reg = reg2; ty = Point2
		// iter 1: reg = reg2; ty = Point3
		// iter 2: reg = reg2; ty = int
		reg = reg2
		ty = strucEntry.Fields[f.ids[i+1]].Ty
	}
	return reg
}

