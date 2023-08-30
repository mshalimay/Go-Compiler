package ir

import (
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)


var createdReadGlobal = false

// call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]*@.read, i32 0, i32 0), <ty>* <op_%d>)
// call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* @globalInit)
type Scanf struct {
	operand 	Operand
}


func NewScanf(operand Operand) *Scanf {
	return &Scanf{operand}
}

func (s *Scanf) String() string {
	if !createdReadGlobal {
		EOFIns.AddInstruction("@.read = private unnamed_addr constant [4 x i8] c\"%ld\\00\", align 1")
		createdReadGlobal = true
	}

	return fmt.Sprintf("call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), %s* %s)\n", 
				s.operand.LLVMTy(), s.operand.String())
}


func (s *Scanf) GetTargets() []int{
	return nil
}


func (s *Scanf) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	code := codegen.Instructions{}
	code.WriteComment(s.String())

	reg := codegen.GetFreeRegister()
	code.WriteCode(fmt.Sprintf("adrp %s, .READ", reg.String))
	code.WriteCode(fmt.Sprintf("add %s, %s, :lo12:.READ", reg.String, reg.String))
	code.WriteCode(fmt.Sprintf("mov x0, %s", reg.String))
	
	block, location := codegen.GetVariableLocation(s.operand.String()[1:], funCallee)

	// TODO check if global is correct
	if block != "" {
		// if global
		code.WriteCode(fmt.Sprintf("add x1, %s, #0", location))
	} else {
		// if local
		location := funCallee.StackRegisters[s.operand.String()[1:]]
		// if local
		code.WriteCode(fmt.Sprintf("add x1, sp,  #%s", location))
	}

	code.WriteCode("bl scanf")

	reg.Free()
	return code.String()
}
