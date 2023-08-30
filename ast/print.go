package ast

// inherits from Statement

import (
	"bytes"
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)


type Print struct {
	token 				*token.Token
	str 				string
	expressions 		[]Expression
}

func NewPrint(token *token.Token, str string, expressions []Expression) *Print{
	return &Print{token, str, expressions}
}


func (print *Print) String() string{
	// Print: 'printf' '(' 'string' { ',' Expression} ')'  ';'  

	var out bytes.Buffer
	
	out.WriteString("printf ")
	out.WriteString("(")
	out.WriteString(print.str)

	for _, expression := range print.expressions {
		out.WriteString(", ")
		out.WriteString(expression.String())
	}
	out.WriteString(")")
	out.WriteString(";\n")
	
	return out.String()
}

func (print *Print) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{		
	
	// check if # of "%d" in string matches # of expressions
	placeholderCount := strings.Count(print.str, "%d")
	if placeholderCount != len(print.expressions){
		errors = append(errors, SemErr(print.token,
			fmt.Sprintf("number of placeholders in `printf` does not match number of expressions. Got %d placeholders and %d expressions.",
				placeholderCount, len(print.expressions))))
	}

	for _, expr := range print.expressions {
		errors = expr.TypeCheck(errors, tables, funCallee)
		
		exprType := expr.GetType()
		if exprType != types.IntTySig {
			errors = append(errors, SemErr(print.token, "Expression in `printf` must be of type `int`. Got: "+ 
											exprType.String() + "."))
		}
	}
	return errors
}


func (p *Print) HasReturn(warnings []string) ([]string, bool){
	return warnings, false
}

func (p *Print) GetToken() *token.Token{
	return p.token
}

func (p *Print) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) *ir.BasicBlock {

	
	newStr, charCount := StringToLLVM(p.str)

	// create EOF instruction for printf
	newFstrReg := ir.NewFstrReg()
	eofIns := fmt.Sprintf("%s = private unnamed_addr constant [%d x i8] c\"%s\", align 1", newFstrReg, charCount, newStr)
	
	// create function call to "printf"
	operands := make([]ir.Operand, len(p.expressions))
	
	for i, expr := range p.expressions {
		operands[i] = expr.ToLLVM(tables, funCallee, block)
	}

	block.AddInstruction(ir.NewPrintf(newFstrReg, operands, eofIns, charCount))
	
	return block
}


func StringToLLVM(input string) (string, int) {
	
	input = strings.Trim(input, "\"")
	// Replace %d with %ld
	transformed := strings.Replace(input, "%d", "%ld", -1)

	// Replace \n with \0A
	transformed = strings.Replace(transformed, "\\n", "\\0A", -1)

	// Add \00 at the end
	transformed += "\\00"

	// Compute character count
	charCount := 0
	
	newLineCount := strings.Count(transformed, "\\0A") 

	// %ld = 3 chars; \0A = 1 char; \00 = 1 char

	// charcount = (len(string) - 3 * newLineCount - 3) + newLineCount + 1 (add 1 bcs added \00)
	charCount = len(transformed) - 2 * (newLineCount + 1)

	return transformed, charCount
}