package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

// `Statement` node

type Block struct {
	token      		*token.Token
	statements 		[]Statement
	HasRet 			bool		// used in LLVM conversion to not have to traverse the block twice
}

func NewBlock(token *token.Token, statements []Statement) *Block {
	return &Block{token, statements, false}
}

func (block *Block) String() string {
	// Block = '{' Statements '}'

	var out bytes.Buffer

	out.WriteString("{\n")
	for _, statement := range block.statements {
		out.WriteString(statement.String())
	}
	out.WriteString("}\n")

	return out.String()
}

func (block *Block) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string {
	for _, statement := range block.statements {
		errors = statement.TypeCheck(errors, tables, funCallee)
	}
	return errors
}

func (block *Block) HasReturn(warnings []string) ([]string, bool) {
	hasReturn := false
	var b bool

	for _, statement := range block.statements {
		warnings, b = statement.HasReturn(warnings)
		hasReturn = hasReturn || b

		if hasReturn {
			// REVIEW create a WARNINGS slice for unreachable code
			warnings = append(warnings, SemErr(block.token, "Unreachable code."))
		}
	}
	block.HasRet = hasReturn
	return warnings, hasReturn
}

func (block *Block) GetToken() *token.Token {
	return block.token
}


func (b *Block) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) *ir.BasicBlock {
	curBlock := block
	for _, statement := range b.statements {
		curBlock = statement.ToLLVM(tables, funCallee, curBlock)
	}
	return curBlock
}