package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
	"golite/types"
)

// <result> = call <ty> <fnval>(<ty> <op>, <ty> <op>, ...)
type FunCall struct {
	destReg 	int
	ty 			types.Type
	fnval 		string
	args 		[]Operand
}

func NewFunCall(destReg int, ty types.Type, fnval string, args []Operand) *FunCall {
	return &FunCall{destReg, ty, fnval, args}
}


func (f *FunCall) String() string {
	var out string
	
	if f.destReg != -1 {
		out = fmt.Sprintf("%%r%d = ", f.destReg)
	}

	out += fmt.Sprintf("call %s %s(", types.ToLLVM(f.ty), f.fnval)
	
	for i, arg := range f.args {
		out += fmt.Sprintf("%s %s", arg.LLVMTy(), arg.String())
		if i != len(f.args) - 1 {
			out += ", "
		}
	}

	out += ")\n"
	return out
}


func(f *FunCall) GetTargets() []int {
	return []int{f.destReg}
}


func (f *FunCall) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	code := codegen.Instructions{}
	code.WriteComment(f.String())
	// TODO: complete this
	for i, arg := range f.args {	
		if i < 8 {
			argLoc := funCallee.StackRegisters[arg.Name()]
			code.WriteCode(fmt.Sprintf("ldr x%d, [sp #%s]", i, argLoc))

		} else {
			panic("Not implemented: function call with more than 8 arguments")
		}
	}
	
	retLoc := funCallee.StackRegisters[fmt.Sprintf("r%d", f.destReg)]
	code.WriteCode(fmt.Sprintf("bl %s", f.fnval[1:]))
	code.WriteCode(fmt.Sprintf("str x0, [sp, #%s]", retLoc))

	return code.String()

}