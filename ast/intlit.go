package ast

import (
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

//TODO: remove the int64 (why did I use this?)

// inherits from Expression
type IntLiteral struct {
	token 				*token.Token
	Value 				int64 		// The value of the literal. e.g. 42
	ty 	  				types.Type
}

func NewIntLiteral(value int64, token *token.Token) *IntLiteral {
	return &IntLiteral{token, value, types.IntTySig}
}

func (intlit *IntLiteral) String() string{
	// IntLiteral: [0-9]+
	return fmt.Sprintf("%d", intlit.Value)
}

func (intlit *IntLiteral) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string {
	return errors
}

func (intlit *IntLiteral) GetType() types.Type {
	return intlit.ty
}

func (intlit *IntLiteral) GetToken() *token.Token {
	return intlit.token
}

func (intlit *IntLiteral) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	return ir.NewLiteral(intlit.ty, int(intlit.Value))
}