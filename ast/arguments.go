package ast

import (
	"bytes"
	"fmt"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// NOTE: this node may be deleted from the AST, but helps in modularity

type Arguments struct{
	token 		*token.Token
	expressions []Expression
}


func NewArguments(token *token.Token, expressions []Expression) *Arguments{
	return &Arguments{token, expressions}
}

func (args *Arguments) String() string{
	//arguments: '(' (expression (',' expression)*)? ')';
	
	var out bytes.Buffer
	out.WriteString("(")

	// iterate through expressions and add them to the string
	for i, expression := range args.expressions{
		out.WriteString(expression.String())
		
		// if not last expression, add a COMMA before next expression
		if i != len(args.expressions) - 1{
			out.WriteString(", ")
		}
	}
	out.WriteString(")")

	return out.String() // eg.: (arg1, arg2, arg3)
}


func (args *Arguments) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry, funCalled *st.FuncEntry) []string{

	// get # arguments in signature of called function
	nFunCalled := len(funCalled.Parameters)

	// if num arguments in function signature diff arguments in the invocation, error
	if nFunCalled != len(args.expressions){
		errors = append(errors, SemErr(args.token, 
								fmt.Sprintf("Incorrect number of arguments in function call. Expected %d, got %d.", 
								nFunCalled, len(args.expressions))))
	}
	
	// check if the type of each argument in invocation matches the type in the function signature
	for i, expression := range args.expressions{
		// typecheck the epxressions inside the invocation eg: foo(a, 1+2, a.b.c) => typecheck a, 1+2, a.b.c
		errors = expression.TypeCheck(errors, tables, funCallee)
		
		// get the type of the expression
		exprType := expression.GetType()

		// if exprType nil, expression was not declared; already added to errors in TypeCheck
		// REVIEW throw other errors or skip? Skip unknown types?
		if exprType == types.UndeclTySig {
			continue
		}

		// if i <= nFunCalled, compare types; else, errors: too many arguments
		if i < nFunCalled {
			if exprType != funCalled.Parameters[i].Ty {
				errors = append(errors, SemErr(expression.GetToken(), 
						fmt.Sprintf("Incorrect type of argument in function call of %s. Expected %s, got %s.",
						funCalled.Name, funCalled.Parameters[i].Ty, exprType)))
			}
		} else {
			errors = append(errors, SemErr(expression.GetToken(), 
									"too many arguments in function call of " + funCalled.Name))
		}
	}
	return errors
}


func (args *Arguments) GetToken() *token.Token{
	return args.token
}

		