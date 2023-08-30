package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// inherits from Instruction
// <result> = add <ty> <op1>, <op2>

type Add struct {
	destReg 	int
	op1			Operand
	op2			Operand
}


func NewAdd(destReg int, op1 Operand, op2 Operand) *Add {
	return &Add{destReg, op1, op2}
}


func (a *Add) String() string {
	return fmt.Sprintf("%%r%d = add %s %s, %s\n", a.destReg, a.op1.LLVMTy(), a.op1.String(), a.op2.String())
}


func (a *Add) GetTargets() []int {
	return []int{a.destReg}
}

func (a *Add) Op1() Operand {
	return a.op1
}

func (a *Add) Op2() Operand {
	return a.op2
}


func (ins *Add) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	
	code := codegen.Instructions{}
	code.WriteComment(ins.String())
	
	// get free registers
	var lReg, rReg, codeBlock string
	var r1, r2 *codegen.ARMRegister
	
	// `op1` can come from literal, global, register or stack		
	lReg, codeBlock, r1 = GetRegister(ins, ins.op1, *funCallee)
	code.WriteBlock(codeBlock)

	// `op2` can come from literal, global, register or stack		
	rReg, codeBlock, r2 = GetRegister(ins, ins.op2, *funCallee)
	code.WriteBlock(codeBlock)
	
	// add r3, r2, r1
	reg := codegen.GetFreeRegister()
	code.WriteCode(fmt.Sprintf("add %s, %s, %s", reg.String, lReg, rReg))

	// store result to target register
	target := funCallee.StackRegisters[fmt.Sprintf("r%d", ins.destReg)]

	if target[0] == 'x'{
		// mov <dest>, <source>
		code.WriteCode(fmt.Sprintf("mov %s, %s", target, reg.String))
	} else {
		// str <source>, [<destination>]
		code.WriteCode(fmt.Sprintf("str %v, [sp, #%v]", reg.String, target))
	}

	reg.Free()
	r1.Free()
	r2.Free()

	return code.String()
}


