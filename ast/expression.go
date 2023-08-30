package ast

import (
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Operator int

const (
	ADD Operator = iota
	SUB
	MUL
	DIV
	NOT
	GT
	GTE
	LT
	LTE
	EQ
	NEQ
	AND
	OR
)

// StrToOperator converts a string to an operator
// e.g. "+" -> ADD
// Obs: if string is empty, returns -1
func StrToOperator(op string) Operator{
	switch op {
	case "+":
		return ADD
	case "-":
		return SUB
	case "*":
		return MUL
	case "/":
		return DIV
	case "!":
		return NOT
	case ">":
		return GT
	case ">=":
		return GTE
	case "<":
		return LT
	case "<=":
		return LTE
	case "==":
		return EQ
	case "!=":
		return NEQ
	case "&&":
		return AND
	case "||":
		return OR
	case "":
		return -1
	}
	panic("Error: Could not determine operator")
}

// OpToStr converts an operator to a string
// e.g. ADD -> "+"
// Obs: if operator is -1, returns empty string
func OpToStr(op Operator) string {
	switch op {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case NOT:
		return "!"
	case GT:
		return ">"
	case GTE:
		return ">="
	case LT:
		return "<"
	case LTE:
		return "<="
	case EQ:
		return "=="
	case NEQ:
		return "!="
	case AND:
		return "&&"
	case OR:
		return "||"
	case -1:
		return ""
	}
	panic("Error: Could not determine operator")
}

func IsRelational(op Operator) bool {
	return op == GT || op == GTE || op == LT || op == LTE
}

func IsLogical(op Operator) bool {
	return op == AND || op == OR
}

func IsArithmetic(op Operator) bool {
	return op == ADD || op == SUB || op == MUL || op == DIV
}


func OpToLLVM(op Operator) string{
	switch op {
	case ADD:
		return "add"
	case SUB:
		return "sub"
	case MUL:
		return "mul"
	case DIV:
		return "sdiv"
	case NOT:
		return "xor"
	case GT:
		return "sgt"
	case GTE:
		return "sge"
	case LT:
		return "slt"
	case LTE:
		return "sle"
	case EQ:
		return "eq"
	case NEQ:
		return "ne"
	case AND:
		return "and"
	case OR:
		return "or"
	}
	panic("Error: Could not determine operator")
}


// Expression is the base node for all expression specific nodes (i.e., every expression sub node implements this
// interface)
type Expression interface {
	String() 	string
	GetToken() 	*token.Token
	GetType() 	types.Type
	TypeCheck([]string, *st.SymbolTables, *st.FuncEntry) 	[]string
	ToLLVM(*st.SymbolTables, *st.FuncEntry, *ir.BasicBlock) ir.Operand
}

