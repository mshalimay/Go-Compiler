package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// ret <type> <value>
type Ret struct {
	value  Operand
}


func NewRet(value Operand) *Ret{
	return &Ret{value}
}

func (ret *Ret) String() string {
	if ret.value != nil {
		return fmt.Sprintf("ret %s %s\n", ret.value.LLVMTy(), ret.value.String())
	
	} else {
		return "ret void\n"
	}
}


func (ret *Ret) GetTargets() []int{
	return nil
}


func (ret *Ret) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	code := codegen.Instructions{}
	code.WriteComment(ret.String())

	if ret.value != nil {
		code.WriteCode(fmt.Sprintf("ldr x0, [sp, #%s]", ret.value.String()))
	}
	return code.String()
}

