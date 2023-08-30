package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)


type LValue struct{
	token 		*token.Token
	variable 	Expression		// *ast.Variable 
	field 		Expression		// *ast.Field
	ty          types.Type		
}


func NewLValue(token *token.Token, variable Expression, field Expression, ty types.Type) *LValue{
	return &LValue{token, variable, field, nil}
}


func (lvalue *LValue) String() string{
	// LValue: 'id' {'.' id}
	var out bytes.Buffer

	if lvalue.variable != nil {
		out.WriteString(lvalue.variable.String())
	} else {
		out.WriteString(lvalue.field.String())
	}
	return out.String()
}


func (lvalue *LValue) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{
	if lvalue.variable != nil {
		errors = lvalue.variable.TypeCheck(errors, tables, funCallee)
		lvalue.ty = lvalue.variable.GetType()
	} else {
		errors = lvalue.field.TypeCheck(errors, tables, funCallee)
		lvalue.ty = lvalue.field.GetType()
	}
	return errors
}


func (lvalue *LValue) GetType() types.Type {
	return lvalue.ty
}

func (lvalue *LValue) GetToken() *token.Token {
	return lvalue.token
}

func (lvalue *LValue) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	
	var operand ir.Operand

	// if lvalue is parameter, local, or global variable, get `operand` from symbol table
	// 	 obs: compute here because ToLLVM method of `variable` loads to memory for the case variable in the RHS of an assignment
	if lvalue.variable != nil {
		// get varEntry from symbol table
		v := lvalue.variable.(*Variable)
		varEntry := GetVarEntry(tables, funCallee, v.id)
		operand = varEntry.Reg.(ir.Operand)

	// if lvalue is field, call ToLLVM method of `field` to get Operand = %r<label>
	} else {
		operand = lvalue.field.(*Field).ToLLVMAux(tables, funCallee, block)
	}

	return operand
}