package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

// `Statement` node
  // obs: `Invocation` does not appear in the RHS of assignment
  // it is a call to a function amidst the source code.
  // eg: foo(1, 2, 3);  	// NOT: a = foo(1, 2, 3);
type Invocation struct {
	token 			*token.Token
	id 				string
	arguments 		*Arguments
}


func NewInvocation(token *token.Token, id string, arguments *Arguments) *Invocation{
	return &Invocation{token, id, arguments}
}

func (invocation *Invocation) String() string{
	// Invocation: 'id' Arguments ';'
	// eg: foo(1, 2, 3);

	var out bytes.Buffer

	out.WriteString(invocation.id)
	out.WriteString(invocation.arguments.String())
	out.WriteString(";\n")
	
	return out.String()
}


func (invoc *Invocation) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{
	// check if the invocated function exists
	funCalled, ok := tables.Funcs.Contains(invoc.id)

	if !ok {
		errors = append(errors, SemErr(invoc.token, "call to non-declared function `"+invoc.id+"`."))
		return errors
	}
	errors = invoc.arguments.TypeCheck(errors, tables, funCallee, funCalled)
	return errors
}


func (invoc *Invocation) HasReturn(warnings []string) ([]string, bool){
	return warnings, false
}

func (invoc *Invocation) GetToken() *token.Token{
	return invoc.token
}


func (invoc *Invocation) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) *ir.BasicBlock {

	// load arguments into registers
	regs := make([]ir.Operand, len(invoc.arguments.expressions))
	
	for i, expr := range invoc.arguments.expressions {
		regs[i] = expr.ToLLVM(tables, funCallee, block)
	}

	// add instruction for function call
	funcEntry, _ := tables.Funcs.Contains(invoc.id)
	
	block.AddInstruction(ir.NewFunCall(-1, funcEntry.RetTy, funcEntry.LLVMName, regs))
	return block
}
