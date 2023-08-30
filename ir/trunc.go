package ir

import (
	"fmt"
	st "golite/symboltable"
)

// <result> = trunc <ty> <value> to <destty>
type Trunc struct {
	toReg 		int
	tyFrom 		string
	fromReg 	Operand
	tyTo 		string
}

func NewTrunc(toReg int, tyFrom string, fromReg Operand, tyTo string) *Trunc {
	return &Trunc{toReg, tyFrom, fromReg, tyTo}
}


func (t *Trunc) String() string {
	return fmt.Sprintf("%%r%d = trunc %s %s to %s\n", t.toReg, t.tyFrom, t.fromReg.String(), t.tyTo)
}


func (t *Trunc) GetTargets() []int {
	return []int{t.toReg}
}

func (t *Trunc) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	return ""
}