package ir

import st "golite/symboltable"

// REVIEW evaluate if it is worth to combine malloc, printf, scan, free single in FunCall struct

// call void @free(i8* %r23)

type Free struct {
	op 	Operand
}

func NewFree(regFree Operand) *Free {
	return &Free{regFree}
}

func (f *Free) String() string {
	return "call void @free(i8* " + f.op.String() + ")\n"
}


func(f *Free) GetTargets() []int {
	return nil
}

func (f *Free) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	// TODO IMPLEMENT
	return ""
}