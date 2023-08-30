package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)



type Instruction interface {
	/* Get the registers targeted by this instruction */
	GetTargets() []int
	// /* Get the source registers for this instruction */
	// GetSources() []int
	// /* Get the immediate value of this instruction */
	// GetImmediate() int
	// /* Get the label that marks this instruction in code */
	// GetLabel() int
	// /* Set the label that marks this instruction in code */
	// SetLabel(newLabel int)
	// /* Return a string representation of this instruction */
	String() string
	TranslateToAssembly(*st.SymbolTables, *st.FuncEntry) string
}

type BinopInstruction interface {
	Instruction
	Op1() Operand
	Op2() Operand
}


func GetRegister (ins BinopInstruction, op Operand, funCallee st.FuncEntry) (string, string, *codegen.ARMRegister){
	code := codegen.Instructions{}
	// code.WriteComment(ins.String())

	reg1 := codegen.GetFreeRegister()

	var reg string

	// `op` comes from global
	if _, ok := op.(*GlobalRegister); ok {
		code.WriteCode(fmt.Sprintf("adrp %s, %s", reg1.String, op.Name()))
		code.WriteCode(fmt.Sprintf("add %s, [%s, #:lo12:%s]", reg1.String, reg1.String, op.Name()))
		code.WriteCode(fmt.Sprintf("ldr %s, [%s]", reg1.String, reg1.String))
		reg = reg1.String

	// `op1` comes from literal 			TODO: make this more general 
	} else if lit, ok := op.(*Literal); ok {
		if lit.value < 0 {
			code.WriteCode(fmt.Sprintf("mvn %s, #%d", reg1.String, lit.value-1))
		} else {
			code.WriteCode(fmt.Sprintf("mov %s, #%d", reg1.String, lit.value))
		}
		reg = reg1.String

	// `op1` comes from register or stack
	} else {
		op1Loc := funCallee.StackRegisters[op.Name()]
		if op1Loc[0] == 'x' {
			reg = op1Loc
		} else {
			code.WriteCode(fmt.Sprintf("ldr %s, [sp, #%s]", reg1.String, op1Loc))
			reg = reg1.String
		}
	}

	return reg, code.String(), reg1
}
