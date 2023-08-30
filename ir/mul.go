package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// inherits from Instruction
// <result> = add <ty> <op1>, <op2>

type Mul struct {
	destReg 	int
	op1			Operand
	op2			Operand
}


func NewMul(destReg int, op1, op2 Operand) *Mul {
	return &Mul{destReg, op1, op2}
}


func (m *Mul) String() string {
	return fmt.Sprintf("%%r%d = mul %s %s, %s\n", m.destReg, m.op1.LLVMTy(), m.op1.String(), m.op2.String())
}


func (m *Mul) GetTargets() []int {
	return []int{m.destReg}
}

func (m *Mul) Op1() Operand {
	return m.op1
}

func (m *Mul) Op2() Operand {
	return m.op2
}



func (ins *Mul) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	
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
