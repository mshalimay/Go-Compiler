package ast

import (
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// inherits from Expression
type BoolLiteral struct {
	token 			*token.Token
	Value 			bool
	ty 	  			types.Type	
}

func NewBoolLiteral(value bool, token *token.Token) *BoolLiteral {
	return &BoolLiteral{token, value, types.BoolTySig}
}

func (boolLit *BoolLiteral) String() string{
	// BoolLiteral: 'true' | 'false'
	return fmt.Sprintf("%v", boolLit.Value)
}


func (boolLit *BoolLiteral) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string {
	return errors
}


func (boolLit *BoolLiteral) GetType() types.Type {
	return boolLit.ty
}

func (boolit *BoolLiteral) GetToken() *token.Token {
	return boolit.token
}


func (boolit *BoolLiteral) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	if boolit.Value {
		return ir.NewLiteral(types.IntTySig, 1)
	} else {
		return ir.NewLiteral(types.IntTySig, 0)
	}
}