package ir

import (
	"fmt"
	st "golite/symboltable"
)

// REVIEW evaluate if it is worth to combine malloc, printf, scan single in FunCall struct

// <result> = call i8* @malloc(i32 <size>)
type Malloc struct {
	destReg 	int
	size 		int
}

func NewMalloc(destReg int, size int) *Malloc {
	return &Malloc{destReg, size}
}

func (m *Malloc) String() string {
	return fmt.Sprintf("%%r%d = call i8* @malloc(i32 %d)\n", m.destReg, m.size)
}

func (m *Malloc) GetTargets() []int {
	return []int{m.destReg}
}


func (m *Malloc) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	return ""
}