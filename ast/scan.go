package ast

import (
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)


//NOTE: at the moment, SCAN does not have antlr.ctx hence, does not have unique key.
// In the parser.nodes map the Token will be using the same key as the `Assignemnt` that contains `Scan`.
type Scan struct {
	token 		*token.Token 
}

func NewScan(token *token.Token) *Scan {
	return &Scan{token}
}

func (scan *Scan) String() string{
	// Scan: 'scan'
	return "scan"
}


func (scan *Scan) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string {
	return errors
}

 func (scan *Scan) GetToken() *token.Token {
	return scan.token
}

func (scan *Scan) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	return nil
}