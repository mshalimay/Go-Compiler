package ast

// `Statement` node

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Delete struct {
	token 			*token.Token
	Expression 		Expression
}

func NewDelete(token *token.Token, expression Expression) *Delete{
	return &Delete{token, expression}
}

func (delete *Delete) String() string{
	// Delete: 'delete' Expression ';'

	var out bytes.Buffer
	out.WriteString("delete ")
	out.WriteString(delete.Expression.String())
	out.WriteString(";\n")
	
	return out.String()
}


func (del *Delete) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{
	errors = del.Expression.TypeCheck(errors, tables, funCallee)
	exprType := del.Expression.GetType()

	if !types.IsStructType(exprType){
		errors = append(errors, SemErr(del.token, "Expression in `delete` must a pointer to a struct. Got: "+
										exprType.String() + "."))
	}
	return errors
}


func (del *Delete) HasReturn(warnings []string) ([]string, bool){
	return warnings, false
}

func (del *Delete) GetToken() *token.Token{
	return del.token
}


func (del *Delete) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) *ir.BasicBlock{


	fromReg := del.Expression.ToLLVM(tables, funCallee, block)
	toReg := ir.NewTempRegister(nil, funCallee)
	block.AddInstruction(ir.NewBitCast(toReg.GetLabel(), fromReg.LLVMTy(), fromReg, "i8*"))

	block.AddInstruction(ir.NewFree(toReg))
	
	return block
}