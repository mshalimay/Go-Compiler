package ir

import (
	"bytes"
	"fmt"
	st "golite/symboltable"
)

var labelCounter int = -1

type BasicBlock struct {
	label 		 int
	instructions []Instruction
	predecessors []*BasicBlock
	successors 	 []*BasicBlock
	BlockTy 	 BlockType
}

func NewBasicBlock(blockTy BlockType) *BasicBlock {
	labelCounter++
	return &BasicBlock{labelCounter, make([]Instruction, 0), make([]*BasicBlock, 0), make([]*BasicBlock, 0), blockTy}
}

func (bb *BasicBlock) AddInstruction(instr Instruction) {
	bb.instructions = append(bb.instructions, instr)
}

func (bb *BasicBlock) GetLastInstruction() Instruction {
	return bb.instructions[len(bb.instructions)-1]
}

func (bb *BasicBlock) GetLabel() int {
	return bb.label
}

func (bb *BasicBlock) String() string {
	var out bytes.Buffer
	
	out.WriteString(fmt.Sprintf("L%d:\n", bb.label))	
	
	for _, ins := range bb.instructions {
		out.WriteString("\t")
		out.WriteString(ins.String())
	}

	for _, succ := range bb.successors {
		out.WriteString(succ.String())
	}
	return out.String()
}


func (bb *BasicBlock) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf(".L%d:\n", bb.label))

	for _, ins := range bb.instructions {
		out.WriteString(ins.TranslateToAssembly(tables, funCallee))
	}
	
	for _, succ := range bb.successors {
		out.WriteString(succ.TranslateToAssembly(tables, funCallee))
	}
	return out.String()
}


func (bb *BasicBlock) AddSuccessor(block *BasicBlock) {
	bb.successors = append(bb.successors, block)
}

func (bb *BasicBlock) AddPredecessor(block *BasicBlock) {
	bb.predecessors = append(bb.predecessors, block)
}

type BlockType int
const (
	BASIC_ENTRY BlockType = iota
	BASIC_EXIT
	IF_ENTRY
	IF_TRUE
	IF_ELSE
	IF_EXIT
	FOR_ENTRY
	FOR_BODY
	FOR_EXIT
)
