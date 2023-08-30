package ast

//NOTE: this node is used for parsing composite term `LPAREN expression RPAREN`
// in the factor production rule.

// `Expression` node

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type FactorPrime1 struct{
	token 			*token.Token
	expression 		Expression
	ty 				types.Type
}

func NewFactorPrime1(token *token.Token, expression Expression, ty types.Type) *FactorPrime1{
	return &FactorPrime1{token, expression, ty}
}

func (fprime *FactorPrime1) String() string{
	// FactorPrime1: '(' Expression ')'
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(fprime.expression.String())
	out.WriteString(")")

	return out.String()
}

func (fprime *FactorPrime1) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{
	// FactorPrime1: '(' Expression ')'
	errors = fprime.expression.TypeCheck(errors, tables, funCallee)
	fprime.ty = fprime.expression.GetType()
	return errors
}


func (fprime *FactorPrime1) GetType() types.Type {
	return fprime.ty
}

func (fprime *FactorPrime1) GetToken() *token.Token {
	return fprime.token
}

func (fprime *FactorPrime1) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	return fprime.expression.ToLLVM(tables, funCallee, block)
}