package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// inherits from Expression
type UnaryTerm struct{
	token 		    *token.Token
	operator 		Operator
	factor 			Expression
	ty 				types.Type
}


 func NewUnaryTerm(token *token.Token, operator Operator, factor Expression, ty types.Type) *UnaryTerm{
	return &UnaryTerm{token, operator, factor, nil}
 }


 func (uTerm *UnaryTerm) String() string{
	// unaryTerm: NOT selectorTerm | MINUS selectorTerm | selectorTerm;
	// selectorTerm: factor selectorTermPrime*;  
	// selectorTermPrime: DOT ID;

	var out bytes.Buffer

	if uTerm.operator != -1{
		out.WriteString(OpToStr(uTerm.operator))
	}
	
	out.WriteString(uTerm.factor.String())	
	return out.String()
 }


 func (uTerm *UnaryTerm) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string {
	// unaryTerm: NOT selectorTerm | MINUS selectorTerm | selectorTerm;

	// typecheck selectorTerm and get it's type
	errors = uTerm.factor.TypeCheck(errors, tables, funCallee)
	factorType := uTerm.factor.GetType()

	// if factor is undeclared => return errors; dont need to return evaluation error
	if factorType == types.UndeclTySig {
		uTerm.ty = types.UndeclTySig
		return errors
	}

	switch uTerm.operator {
		case NOT: // NOT selectorTerm
			if factorType != types.BoolTySig{
				errors = append(errors, SemErr(uTerm.token, " `!` expects a boolean expression. Got: " +
												factorType.String() + "."))
				uTerm.ty = types.UnknownTySig
			} else {
				uTerm.ty = types.BoolTySig
			}
		
		case SUB: // MINUS selectorTerm
			if factorType != types.IntTySig {
				errors = append(errors, SemErr(uTerm.token, " `-` expects an integer expression. Got: " +
												factorType.String() + "."))
				uTerm.ty = types.UnknownTySig
			} else {
				uTerm.ty = types.IntTySig
			}
		
		case -1: // selectorTerm
			uTerm.ty = factorType
		
		default:
			panic("Invalid syntax for unaryTerm")
	}

	return errors
 }


 func (uTerm *UnaryTerm) GetType() types.Type {
	return uTerm.ty
 }

 func (uTerm *UnaryTerm) GetToken() *token.Token {
	return uTerm.token
}


func (uTerm *UnaryTerm) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	// unaryTerm: NOT selectorTerm | MINUS selectorTerm | selectorTerm;
	op1 := uTerm.factor.ToLLVM(tables, funCallee, block)
	
	switch uTerm.operator {
	case NOT:	
		operand := ir.NewTempRegister(uTerm.ty, funCallee)
		block.AddInstruction(ir.NewXor(operand.GetLabel(), op1))
		return operand
	
	case SUB:		
		operand := ir.NewTempRegister(uTerm.ty, funCallee)
		block.AddInstruction(ir.NewSub(operand.GetLabel(), ir.NewLiteral(types.IntTySig, 0), op1))		
		return operand

	default:
		return op1 
	}
}
