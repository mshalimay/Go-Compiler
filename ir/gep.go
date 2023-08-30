package ir

import (
	"fmt"
	st "golite/symboltable"
)

// <dest> = getelementptr <ty>, <ty>* <ptrval>, i32 0, i32 <index>
// %r8 = getelementptr %struct.Point2D, %struct.Point2D* %r7, i32 0, i32 0

type Gep struct {
	destReg 	int
	ptrval		Operand
	index		int
}

func NewGep(destReg int, ptrval Operand, index int) *Gep {
	return &Gep{destReg, ptrval, index}
}


func (g *Gep) String() string {
	strucName := g.ptrval.Ty().String()
	llvty := g.ptrval.LLVMTy()
	return fmt.Sprintf("%%r%d = getelementptr %%struct.%s, %s %s, i32 0, i32 %d\n", g.destReg, strucName, llvty, g.ptrval.String(), g.index)
}


func (g *Gep) GetTargets() []int{
	return []int{g.destReg}
}


func (g *Gep) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	// TODO implement
	return ""
}
