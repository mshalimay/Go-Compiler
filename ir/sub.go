package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// inherits from Instruction
// <result> = add <ty> <op1>, <op2>

type Sub struct {
	destReg 	int
	op1			Operand
	op2			Operand
}


func NewSub(destReg int, op1 Operand, op2 Operand) *Sub {
	return &Sub{destReg, op1, op2}
}


func (s *Sub) String() string {
	return fmt.Sprintf("%%r%d = sub %s %s, %s\n", s.destReg, s.op1.LLVMTy(), s.op1.String(), s.op2.String())
}

func (s *Sub) GetTargets() []int {
	return []int{s.destReg}
}

func (s *Sub) Op1() Operand {
	return s.op1
}

func (s *Sub) Op2() Operand {
	return s.op2
}

func (ins *Sub) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	
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
	
	// sub r3, r2, r1
	reg := codegen.GetFreeRegister()
	code.WriteCode(fmt.Sprintf("sub %s, %s, %s", reg.String, lReg, rReg))

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
