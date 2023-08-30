package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// <result> = xor <ty> <op1>, <op2>

type Xor struct {
	destReg 	int
	op1			Operand
}

func NewXor(destReg int, op1 Operand) *Xor {
	return &Xor{destReg, op1}
}

func (x *Xor) String() string {
	return fmt.Sprintf("%%r%d = xor %s %s, 1\n", x.destReg, x.op1.LLVMTy(), x.op1.String())
}


func (x *Xor) GetTargets() []int {
	return []int{x.destReg}
}

func (x *Xor) Op1() Operand {
	return x.op1
}

func (x *Xor) Op2() Operand {
	return nil
}

func (ins *Xor) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	
	code := codegen.Instructions{}
	code.WriteComment(ins.String())
	
	var lReg, codeBlock string
	var r1 *codegen.ARMRegister

	// `op1` can come from literal, global, register or stack		
	lReg, codeBlock, r1 = GetRegister(ins, ins.op1, *funCallee)
	code.WriteBlock(codeBlock)
	
	// xor r3, r2, r1
	reg := codegen.GetFreeRegister()
	code.WriteCode(fmt.Sprintf("eor %s, %s, %s", reg.String, lReg, "#1"))

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

	return code.String()
}


	
