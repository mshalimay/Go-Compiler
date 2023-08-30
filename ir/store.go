package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
	"golite/types"
)

// inherits from Instruction

// store <ty> <value>, <ty>* <destReg>

type Store struct {
	value 		Operand
	destReg 	Operand
}


func NewStore(value Operand, pointer Operand) *Store {
	return &Store{value, pointer}
}


func (s *Store) String() string {
	ty := s.destReg.LLVMTy()
	return fmt.Sprintf("store %s %s, %s* %s\n", ty, s.value.String(), ty, s.destReg.String())
}


func (s *Store) GetTargets() []int{
	return nil
}

func (s *Store) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	// str, r1, [r2]
	instr := codegen.Instructions{}
	instr.WriteComment(s.String())
	
	// get a free register
	reg1 := codegen.GetFreeRegister()
	reg2 := codegen.GetFreeRegister()

	// 
	var srcReg string
	var sourceIsRegister bool = false

	// `value` can come from literal, global, register or stack
	
	// `value` comes from global
	if _, ok := s.value.(*GlobalRegister); ok {
		instr.WriteCode(fmt.Sprintf("adrp %s, %s", reg1.String, s.value.Name()))
		instr.WriteCode(fmt.Sprintf("str %s, [%s, #:lo12:%s]", reg1.String, reg1.String, s.value.Name()))
		srcReg = reg1.String

	// `value` comes from literal 			TODO: make this more general 
	} else if lit, ok := s.value.(*Literal); ok {
		if lit.value < 0 {
			instr.WriteCode(fmt.Sprintf("mvn %s, #%d", reg1.String, lit.value-1))
		} else {
			instr.WriteCode(fmt.Sprintf("mov %s, #%d", reg1.String, lit.value))
		}
		srcReg = reg1.String

	// register or stack
	} else {
		vLoc := funCallee.StackRegisters[s.value.Name()]
		if vLoc[0] == 'x' {
			srcReg = vLoc
			sourceIsRegister = true
		} else {
			instr.WriteCode(fmt.Sprintf("ldr %s, [sp, #%s]", reg1.String, vLoc))
			srcReg = reg1.String
		}
	}

	// `destReg` can be a global, register or stack
	// `destReg` is a global
	if _, ok := s.destReg.(*GlobalRegister); ok {
		instr.WriteCode(fmt.Sprintf("adrp %s, %s", reg2.String, s.destReg.Name()))
		instr.WriteCode(fmt.Sprintf("str %s, [%s, #:lo12:%s]", reg2.String, reg2.String, s.destReg.Name()))

		// store
		instr.WriteCode(fmt.Sprintf("str %s, [%s]", srcReg, reg2.String))

	// `destReg` is a register or stack
	} else {
		dLoc := funCallee.StackRegisters[s.destReg.Name()]

		if dLoc[0] == 'x' {
			// source and dest are registers => mov
			if sourceIsRegister {
				instr.WriteCode(fmt.Sprintf("mov %s, [sp, #%s]", srcReg, dLoc))
			} else {
				instr.WriteCode(fmt.Sprintf("str %s, [sp, #%s]", srcReg, dLoc))
			}

		} else {
			if types.IsStructType(s.destReg.Ty()) {
				reg3 := codegen.GetFreeRegister()
				instr.WriteCode(fmt.Sprintf("ldr %s, [sp, #%s]", reg3.String, dLoc))
				instr.WriteCode(fmt.Sprintf("str %s, [#%s]", srcReg, reg3.String))
			} else {
				instr.WriteCode(fmt.Sprintf("str %s, [sp, #%s]", srcReg, dLoc))
			}
		}
	}

	reg1.Free()
	reg2.Free()
	return instr.String()
}