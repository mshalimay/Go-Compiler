package ir

import (
	st "golite/symboltable"
)

// Implements Instruction interface

// <destReg> = alloca <ty>
type Alloca struct {
	destReg 		Operand			// LLVM virtual register. This is a pointer to address in stack
}

func NewAlloca(destReg Operand) *Alloca {
	return &Alloca{destReg}
}

func (a *Alloca) String() string {
	return a.destReg.String() + " = alloca " + a.destReg.LLVMTy() + "\n"
}

func (a *Alloca) GetTargets() []int {
	return nil
}




func (a *Alloca) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	return ""
}
