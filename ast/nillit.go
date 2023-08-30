package ast

import (
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// inherits from expression
type NilLit struct {
	token 			*token.Token 
	ty 				types.Type
}


func NewNilLiteral(token *token.Token) *NilLit {
	return &NilLit{token, types.NilTySig}
}

func (nilLit *NilLit) String() string{
	// NilLit: 'nil'
	return "nil"
}

func (nilLit *NilLit) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string {
	return errors
}

func (nilLit *NilLit) GetType() types.Type {
	return nilLit.ty
}

func (nilLit *NilLit) GetToken() *token.Token {
	return nilLit.token
}

func (nilLit *NilLit) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	panic("Called nilLit ToLLVM")
}
