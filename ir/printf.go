package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
	"golite/types"
)

//call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* <printReg>, i32 0, i32 0), <ty> <op>, <ty> <op>, ...)
type Printf struct {
	printReg 		string
	operands 		[]Operand
	eofIns 			string
	strSize			int				
}


func NewPrintf(printReg string, operands []Operand, eofIns string, strSize int) *Printf {
	return &Printf{printReg, operands, eofIns, strSize}
}

func (p *Printf) String() string {
	
	// write EOF instruction for fstring register
	EOFIns.AddInstruction(p.eofIns)

	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([%d x i8], [%d x i8]* %s, i32 0, i32 0), ", 
									p.strSize, p.strSize, p.printReg))
	for i, op := range p.operands {
		out.WriteString(fmt.Sprintf("%s %s", types.ToLLVM(op.Ty()), op.String()))
		if i < len(p.operands) - 1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")\n")

	return out.String()

}


func (p *Printf) GetTargets() []int {
	return nil
}

func (p *Printf) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	code := codegen.Instructions{}
	code.WriteComment(p.String())

	printRegName := p.printReg[2:]

	reg := codegen.GetFreeRegister()
	code.WriteCode(fmt.Sprintf("adrp %s, .FMT%s", reg.String, printRegName))
	code.WriteCode(fmt.Sprintf("add %v, %v, :lo12:.FMT%v", reg.String, reg.String, printRegName))
	code.WriteCode(fmt.Sprintf("mov x0, %s", reg.String))

	// TODO: deal with global variables
	for i, op := range p.operands {
		if _, ok := op.(*Literal); ok {
			code.WriteCode(fmt.Sprintf("mov x%d, %s", i+1, op.String()))
		} else {
			loc := funCallee.StackRegisters[op.Name()]
			code.WriteCode(fmt.Sprintf("ldr x%d, [sp, #%v]", i+1, loc))
		}
	}
	code.WriteCode("bl printf")

	// write EOF instruction
	EOFIns.AddArmInstruction(fmt.Sprintf(".FMT%s:\n", printRegName))
	EOFIns.AddArmInstruction(fmt.Sprintf("\t .asciz \"%%ld\\00\"\n"))
	EOFIns.AddArmInstruction(fmt.Sprintf("\t .size .FMT%s, %d\n", printRegName, p.strSize))


	reg.Free()
	return code.String()

}

