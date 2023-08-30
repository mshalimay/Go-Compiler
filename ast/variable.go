package ast

// inherits from Expression

import (
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// inherits from Expression
type Variable struct {
	token 	   *token.Token
	id 		   string 			// The variable name. e.g. "x"
	ty 		   types.Type 		
}

func NewVariable(id string, token *token.Token, ty types.Type) *Variable {
	return &Variable{token, id, ty}
}

func (v *Variable) String() string { 
	return v.id 
	// return "Var string method: " + v.Identifier 
}

func (v *Variable) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{
	
	// check if is parameter
	if parEntry, contains := funCallee.ContainsParam(v.id); contains{
		v.ty = parEntry.Ty

	// check if is variable
	} else if varEntry, contains := funCallee.Variables.Contains(v.id); contains{
		v.ty = varEntry.Ty

	} else {
		errors = append(errors, SemErr(v.token, "variable " + v.id +" not declared."))
		v.ty = types.UndeclTySig
	}
	return errors
}


func (v *Variable) GetType() types.Type{
	return v.ty
}


func (v *Variable) GetToken() *token.Token{
	return v.token
}


func (v *Variable) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {

	varEntry := GetVarEntry(tables, funCallee, v.id)
	pointer := varEntry.Reg.(ir.Operand)

	// Create new temporary register
	tReg := ir.NewTempRegister(v.ty, funCallee)
	
	// `load` variable value into temp register
	// <dest> = load <ty>, <ty>* <origin>
	block.AddInstruction(ir.NewLoad(tReg.GetLabel(), pointer))
	
	// return temp register
	return tReg
}