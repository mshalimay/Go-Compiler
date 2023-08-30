package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
	"golite/types"
)

// inherits from Instruction

// <result> = load <ty>, <ty>* <pointer>

type Load struct {
	resultReg   int
	pointer 	Operand
}


func NewLoad(resultReg int, pointer Operand) *Load {
	return &Load{resultReg, pointer}
}


func (l *Load) String() string {
	ty := l.pointer.LLVMTy()
	return fmt.Sprintf("%%r%d = load %s, %s* %s\n", l.resultReg, ty, ty, l.pointer.String())
}

func (l *Load) GetTargets() []int {
	return []int{l.resultReg}
}

func (l *Load) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	instrs := codegen.Instructions{}
	instrs.WriteComment(l.String())

	reg := codegen.GetFreeRegister()

	var fromR string

	if _, ok := l.pointer.(*GlobalRegister); ok {
		instrs.WriteCode(fmt.Sprintf("adrp %s, %s", reg.String, l.pointer.Name()))
		instrs.WriteCode(fmt.Sprintf("add %s, %s, #:lo12:%s", reg.String, reg.String, l.pointer.Name()))
		instrs.WriteCode(fmt.Sprintf("ldr %s, [%s]", reg.String, reg.String))

		fromR = reg.String
	} else {
		pLoc := funCallee.StackRegisters[l.pointer.Name()]
		if pLoc[0] == 'x' {
			fromR = pLoc
		} else {
			instrs.WriteCode(fmt.Sprintf("ldr %s, [sp, #%s]", reg.String, pLoc))
			if types.IsStructType(l.pointer.Ty()) {
				instrs.WriteCode(fmt.Sprintf("ldr %s, [%s]", reg.String, reg.String))
			}
			fromR = reg.String
		}
	}

	destLoc := funCallee.StackRegisters[fmt.Sprintf("r%d", l.resultReg)]
	if destLoc[0] == 'x' {
		instrs.WriteCode(fmt.Sprintf("mv %s, %s", destLoc, fromR))
	
	} else {
		instrs.WriteCode(fmt.Sprintf("str %s, [sp, #%s]", fromR, destLoc))
	}

	reg.Free()
	return instrs.String()
}





