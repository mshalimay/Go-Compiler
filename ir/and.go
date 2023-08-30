package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// <result> = and <ty> <op1>, <op2>
type And struct {
	destReg 	int
	op1			Operand
	op2			Operand
}


func NewAnd(destReg int, op1 Operand, op2 Operand) *And {
	return &And{destReg, op1, op2}
}

func (a *And) String() string {
	return fmt.Sprintf("%%r%d = and %s %s, %s\n", a.destReg, a.op1.LLVMTy(), a.op1.String(), a.op2.String())
}


func (a *And) GetTargets() []int {
	return []int{a.destReg}
}

func (a *And) Op1() Operand {
	return a.op1
}

func (a *And) Op2() Operand {
	return a.op2
}



func (ins *And) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	
	code := codegen.Instructions{}
	code.WriteComment(ins.String())
	
	var lReg, rReg, codeBlock string
	
	var r1, r2 *codegen.ARMRegister

	// `op1` can come from literal, global, register or stack		
	lReg, codeBlock, r1 = GetRegister(ins, ins.op1, *funCallee)
	code.WriteBlock(codeBlock)

	// `op2` can come from literal, global, register or stack		
	rReg, codeBlock, r2 = GetRegister(ins, ins.op2,  *funCallee)
	code.WriteBlock(codeBlock)
	
	// and r3, r2, r1
	reg := codegen.GetFreeRegister()
	code.WriteCode(fmt.Sprintf("and %s, %s, %s", reg.String, lReg, rReg))

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


