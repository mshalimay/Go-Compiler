package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// <dest> = icmp <cond> <ty> <op1>, <op2>

type Icmp struct {
	destReg  	int
	cond 		string
	op1 		Operand
	op2 		Operand
}


func NewIcmp(destReg int, cond string, op1, op2 Operand) *Icmp{
	return &Icmp{destReg, cond, op1, op2}
}


func (i *Icmp) String() string {
	return fmt.Sprintf("%%r%d = icmp %s %s %s, %s\n", i.destReg, i.cond, i.op1.LLVMTy(), i.op1.String(), i.op2.String())
}

func (i *Icmp) GetTargets() []int {
	return []int{i.destReg}
}

func (icmp *Icmp) Op1() Operand {
	return icmp.op1
}

func (icmp *Icmp) Op2() Operand {
	return icmp.op2
}


func (ins *Icmp) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
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
	
	// icmp r2, r1
	code.WriteCode(fmt.Sprintf("cmp %s, %s", lReg, rReg))

	// partial branch instruction
	armCond := condToAssembly(ins.cond)
	code.WritePartial("\tb." + armCond + " ")
	
	r1.Free()
	r2.Free()
	return code.String()
}



func condToAssemblyInverse(cond string) string {
	switch cond {
	case "eq":
		return "ne"
	case "ne":
		return "eq"
	case "sgt":
		return "le"
	case "sge":
		return "lt"
	case "slt":
		return "ge"
	case "sle":
		return "gt"
	default:
		panic("invalid icmp condition")
	}
}

func condToAssembly(cond string) string {
	switch cond {
	case "eq":
		return "eq"
	case "ne":
		return "ne"
	case "sgt":
		return "gt"
	case "sge":
		return "ge"
	case "slt":
		return "lt"
	case "sle":
		return "le"
	default:
		panic("invalid icmp condition")
	}
}