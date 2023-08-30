package ast

import (
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

// Statement is the interface for all statements.
type Statement interface {
	String() string
	GetToken() *token.Token
	TypeCheck([]string, *st.SymbolTables, *st.FuncEntry) []string
	HasReturn([]string) ([]string, bool)
	ToLLVM(*st.SymbolTables, *st.FuncEntry, *ir.BasicBlock) *ir.BasicBlock
}


