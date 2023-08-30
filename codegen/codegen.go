package codegen

import (
	"bytes"
	"fmt"
	st "golite/symboltable"
)


// Instructions wrapper
type Instructions struct {
	out 		bytes.Buffer
}

func (ins *Instructions) WriteCode(str string){
	ins.out.WriteString("\t" + str +  "\n")
}

func (ins *Instructions) WritePartial(str string){
	ins.out.WriteString(str)
}

func (ins *Instructions) WriteBlock(str string){
	ins.out.WriteString(str)
}

func (ins *Instructions) WriteComment(str string){
	ins.out.WriteString("//" + str + "\n")
}

func (ins *Instructions) WriteLabel(str string) {
	ins.out.WriteString(str + ":\n")
}

func (ins *Instructions) String() string{
	return ins.out.String()
}

// Register counters
type ARMRegister struct {
	Num 		int
	String 		string
}

var Registers map[int]bool

func InitRegisters() {
	Registers = map[int]bool{}
	for i := 8; i <= 15; i++ {
		Registers[i] = true
	}
}

func GetFreeRegister() *ARMRegister {
	for i :=8; i <= 15; i++{
		if Registers[i] {
			Registers[i] = false
			return &ARMRegister{i, fmt.Sprintf("x%d", i)}
		}
	}
	return &ARMRegister{-1, "not enough registers"}
}


func (reg *ARMRegister) Free() {
	Registers[reg.Num] = true
}



func init() {
	InitRegisters()
}




func GetVariableLocation(varName string, funCallee *st.FuncEntry) (string, string) {
	// search for local variable or parameter
	if vLoc, ok := funCallee.StackRegisters[varName]; ok {
		return "", fmt.Sprintf("[sp, #%s]", vLoc)
	}

	// not in symbol table => global variable	
	instr := Instructions{}
	reg := GetFreeRegister()
	
	// adrp target_register, global_var_name
	instr.WriteCode(fmt.Sprintf("adrp %s, %s", reg.String, varName))

	// add target_register, target_register, :lo12:global_var_name
	instr.WriteCode(fmt.Sprintf("add %s, %s, :lo12:%s", reg.String, reg.String, varName))
	
	// free register
	reg.Free()
	
	return instr.String(), fmt.Sprintf("[%s]", reg.String)
}




