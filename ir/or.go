package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// <result> = or <ty> <op1>, <op2>
type Or struct {
	destReg 	int
	op1			Operand
	op2			Operand
}


func NewOr(destReg int, op1 Operand, op2 Operand) *And {
	return &And{destReg, op1, op2}
}

func (or *Or) String() string {
	return fmt.Sprintf("%%r%d = and %s %s, %s\n", or.destReg, or.op1.LLVMTy(), or.op1.String(), or.op2.String())
}


func (or *Or) GetTargets() []int {
	return []int{or.destReg}
}

func (or *Or) Op1() Operand {
	return or.op1
}

func (or *Or) Op2() Operand {
	return or.op2
}

func (ins *Or) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	
	code := codegen.Instructions{}
	code.WriteComment(ins.String())
	
	var lReg, rReg, codeBlock string
	var r1, r2 *codegen.ARMRegister

	// `op1` can come from literal, global, register or stack		
	lReg, codeBlock, r1 = GetRegister(ins, ins.op1, *funCallee)
	code.WriteBlock(codeBlock)

	// `op2` can come from literal, global, register or stack		
	rReg, codeBlock, r2 = GetRegister(ins, ins.op2, *funCallee)
	code.WriteBlock(codeBlock)
	
	// mul r3, r2, r1
	reg := codegen.GetFreeRegister()
	code.WriteCode(fmt.Sprintf("mul %s, %s, %s", reg.String, lReg, rReg))

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
