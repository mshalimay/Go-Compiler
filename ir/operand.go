package ir

import (
	"fmt"
	st "golite/symboltable"
	"golite/types"
	"strconv"
)


var tempRegLabel = -1
var printRegLabel = 0


func NewTempRegLabel() int {
	tempRegLabel++
	return tempRegLabel
}

func NewFstrReg() string {
	printRegLabel++
	return "@.fstr" + strconv.Itoa(printRegLabel)
}


type Operand interface {
	String() 	string
	Ty() 	 	types.Type
	LLVMTy() 	string
	Name() 		string
}


type Literal struct {
	ty 		types.Type
	value 	int
}


type VirtualRegister struct {
	ty 		types.Type
	name	string
}

type GlobalRegister struct {
	ty 		types.Type
	name	string
}

type TempRegister struct {
	ty 		types.Type
	label 	int
	name 	string
}

func (temp *TempRegister) GetLabel() int {
	return temp.label
}

func NewLiteral(ty types.Type, value int) *Literal {
	return &Literal{ty, value}
}

func NewVirtualRegister(ty types.Type, name string) *VirtualRegister {
	return &VirtualRegister{ty, name}
}

func NewGlobalRegister(ty types.Type, name string) *GlobalRegister {
	return &GlobalRegister{ty, name}
}

func NewTempRegister(ty types.Type, funcEntry *st.FuncEntry) *TempRegister {
	tempRegLabel++ 
	name := fmt.Sprintf("r%d", tempRegLabel)
	funcEntry.TempRegisters = append(funcEntry.TempRegisters, name)
	return &TempRegister{ty, tempRegLabel, name}
}

// String methods

func (l *Literal) String() string {
	return strconv.Itoa(l.value)
}

func (v *VirtualRegister) String() string {
	return "%" + v.name
}

func (g *GlobalRegister) String() string {
	return "@" + g.name
}

func (t *TempRegister) String() string {
	return "%" +  t.name
}

// Ty() methods

func (l *Literal) Ty() types.Type {
	return l.ty
}

func (v *VirtualRegister) Ty() types.Type {
	return v.ty
}

func (g *GlobalRegister) Ty() types.Type {
	return g.ty
}

func (t *TempRegister) Ty() types.Type {
	return t.ty
}

// LLVMTy() methods

func (l *Literal) LLVMTy() string {
	return types.ToLLVM(l.ty)
}

func (v *VirtualRegister) LLVMTy() string {
	return types.ToLLVM(v.ty)
}

func (g *GlobalRegister) LLVMTy() string {
	return types.ToLLVM(g.ty)
}

func (t *TempRegister) LLVMTy() string {
	return types.ToLLVM(t.ty)
}


// Name() methods
func (l *Literal) Name() string {
	return strconv.Itoa(l.value)
}

func (v *VirtualRegister) Name() string {
	return v.name
}

func (g *GlobalRegister) Name() string {
	return g.name
}

func (t *TempRegister) Name() string {
	return t.name
}

