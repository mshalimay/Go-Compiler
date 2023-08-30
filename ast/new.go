package ast

//NOTE: this node parses the composite term `NEW ID` in the factor production rule.

import (
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

//REVIEW USE string or variable for NEW ID? both work, depend on next phases of compiler
type New struct{
	token 		*token.Token
	id 			string
	ty 		   	types.Type
}

func NewNew(token *token.Token, id string, ty types.Type) *New{
	return &New{token, id, ty}
}

func (nw *New) String() string{
	// New: 'new' 'id'
	return "new" + " " + nw.id
}


func (new *New) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{
	// check if identifier is a struct type
	strucEntry, ok := tables.Structs.Contains(new.id)

	// if found struct, assign type to this expression; otherwise, unknown type
	if !ok {
		errors = append(errors, SemErr(new.token, new.id + " is not a struct or is not declared."))
		new.ty = types.UnknownTySig
	} else {
		new.ty = strucEntry.Ty
	}
	return errors
}

func (new *New) GetType() types.Type {
	return new.ty
}

func (new *New) GetToken() *token.Token {
	return new.token
}



func (new *New) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {

	strucEntry, _ := tables.Structs.Contains(new.id)

	// compute memory necessary for malloc
	memSize := len(strucEntry.Fields) * 8

	// call malloc and put result in a register
	// <result> = call i8* @malloc(i32 <size>)

	fromReg := ir.NewTempRegister(nil, funCallee)
	block.AddInstruction(ir.NewMalloc(fromReg.GetLabel(), memSize))

	// cast malloc result to struct type
	// <result> = bitcast <ty> <value> to <ty2>

	toReg := ir.NewTempRegister(strucEntry.Ty, funCallee)
	block.AddInstruction(ir.NewBitCast(toReg.GetLabel(),"i8*", fromReg, toReg.LLVMTy()))
	return toReg
}
