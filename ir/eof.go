package ir

import "bytes"

var EOFIns eofInstructions

type eofInstructions struct {
	instructions 	[]string
	armInstructions []string
}


func (eof *eofInstructions) AddInstruction(ins string) {
	eof.instructions = append(eof.instructions, ins)
}

func (eof *eofInstructions) AddArmInstruction(ins string) {
	eof.armInstructions = append(eof.armInstructions, ins)
}



func (eof *eofInstructions) String() string {
	var out bytes.Buffer

	out.WriteString("declare i8* @malloc(i32)\n")
	out.WriteString("declare void @free(i8*)\n")
	out.WriteString("declare i32 @printf(i8*, ...)\n")
	out.WriteString("declare i32 @scanf(i8*, ...)\n")

	for _, ins := range eof.instructions {
		out.WriteString("\n")
		out.WriteString(ins)
	}
	return out.String()
}


func (eof *eofInstructions) ARMString() string {
	var out bytes.Buffer

	out.WriteString(".READ:\n")
	out.WriteString("\t .asciz \"%ld\\00\"\n")
	out.WriteString("\t .size .READ, 4\n")

	for _, ins := range eof.armInstructions {
		out.WriteString(ins)
	}
	return out.String()
}