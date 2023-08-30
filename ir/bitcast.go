package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// <result> = bitcast <ty> <value> to <ty2>
type BitCast struct {
	toReg 		int
	tyFrom 		string
	fromReg 	Operand
	tyTo 		string
}

func NewBitCast(toReg int, tyFrom string, fromReg Operand, tyTo string) *BitCast {
	return &BitCast{toReg, tyFrom, fromReg, tyTo}
}


func (b *BitCast) String() string {
	return fmt.Sprintf("%%r%d = bitcast %s %s to %s\n", b.toReg, b.tyFrom, b.fromReg.String(), b.tyTo)
}


func (b *BitCast) GetTargets() []int {
	return []int{b.toReg}
}


func (b *BitCast) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	instr := codegen.Instructions{}
	instr.WriteComment(b.String())
	
	// get free registers
	reg1 := codegen.GetFreeRegister()
	reg2 := codegen.GetFreeRegister()

	var fromReg string
	
	// fromReg can come from register or stack
	fLoc := funCallee.StackRegisters[b.fromReg.Name()]
	if fLoc[0] == 'x' {
		fromReg = fLoc
	} else {
		instr.WriteCode(fmt.Sprintf("ldr %s, [sp, #%s]", reg1.String, fLoc))
		fromReg = reg1.String
	}

	// `toReg` can be register or stack	
	toRegStr := fmt.Sprintf("r%d", b.toReg)
	destReg := funCallee.StackRegisters[toRegStr]
	if destReg[0] == 'x' {
		instr.WriteCode(fmt.Sprintf("mov %s, %s", destReg, fromReg))
	} else {
		instr.WriteCode(fmt.Sprintf("ldr %s, [sp, #%s]", reg2.String, fromReg))
		instr.WriteCode(fmt.Sprintf("str %s, [sp, #%s]", reg2.String, destReg))
	}

	reg1.Free()
	reg2.Free()

	return instr.String()
}

