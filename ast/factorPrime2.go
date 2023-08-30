package ast

//NOTE: this node is used for parsing composite term `ID (arguments)?`
// in the factor production rule.

// factorPrime2 is a function call. It is used in the rhs of expressions. It is different from an invocation
// usage examples:
// `a = foo();` ==>  factorPrime2 is foo()
//  `a = foo(1, 2, 3)` ; `c = bar(a,b) + d`;

// `Expression` node

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type FactorPrime2 struct{
	token 			*token.Token
	Identifier 		string 
	Arguments 		*Arguments
	ty 				types.Type
}

func NewFactorPrime2(token *token.Token, identifier string, arguments *Arguments, ty types.Type) *FactorPrime2{
	return &FactorPrime2{token, identifier, arguments, ty}
}

func (fprime *FactorPrime2) String() string{
	// FactorPrime2: ID (Arguments)?

	// NOTE: this will only be called for the case ID(Arguments)
	// for case ID, `variable` string method will be called instead.

	var out bytes.Buffer

	out.WriteString(fprime.Identifier)
	out.WriteString(fprime.Arguments.String())
	
	return out.String()
}

func (fprime *FactorPrime2) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{
	// NOTE: this will only be called for the case ID(Arguments) eg: a = foo(1, 2, 3) + 3; 
	// fPrime2 appears in assignments, conditions, etc; this is an `expression`, not an `invocation`

	// get and check if the called function exists
	funCalled, ok := tables.Funcs.Contains(fprime.Identifier)
	if !ok {
		errors = append(errors, SemErr(fprime.token, "call to non-declared function `"+fprime.Identifier+"`."))
		fprime.ty = types.UndeclTySig
		return errors
	}

	// check if number and types of arguments passed to the called function matches the function signature
	// and set the type of the factorPrime2 node.
	initLen := len(errors)
	errors = fprime.Arguments.TypeCheck(errors, tables, funCallee, funCalled)
	
	if initLen == len(errors) {
		fprime.ty = funCalled.RetTy
	} else {
		fprime.ty = types.UnknownTySig
	}

	return errors
}

func (fprime *FactorPrime2) GetType() types.Type {
	return fprime.ty 
}

func (fprime *FactorPrime2) GetToken() *token.Token{
	return fprime.token
}


func (fprime *FactorPrime2) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	// load arguments into registers
	regs := make([]ir.Operand, len(fprime.Arguments.expressions))
	
	for i, expr := range fprime.Arguments.expressions {
		regs[i] = expr.ToLLVM(tables, funCallee, block)
	}

	// get func from ST and add instruction for function call
	funcEntry, _ := tables.Funcs.Contains(fprime.Identifier)
	
	destReg := ir.NewTempRegister(funcEntry.RetTy, funCallee)
	block.AddInstruction(ir.NewFunCall(destReg.GetLabel(), destReg.Ty(), funcEntry.LLVMName, regs))
	
	return destReg
}