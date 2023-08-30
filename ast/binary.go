package ast

import (
	"bytes"
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

// inherits from Expression
type BinOpExpr struct {
	token 		 *token.Token       // The token information
	operator     Operator   		// The operator for the binary expression
	left         Expression 		// The left operand expression
	right        Expression 		// The right operand expression
	ty		   	 types.Type 		// type of the whole binop expression
}

func NewBinOp(op Operator, left, right Expression, token *token.Token, ty types.Type) *BinOpExpr {
	return &BinOpExpr{token, op, left, right, ty}
}

func (binOp *BinOpExpr) String() string {
	// BinOpExpr: Expression (Operator Expression)*

	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(binOp.left.String())
	out.WriteString(" " + OpToStr(binOp.operator) + " ")
	out.WriteString(binOp.right.String())
	out.WriteString(")")

	return out.String()
}

// All arithmetic and relational operators require integer operands.

// Equality operators require operands of integer or struct type. 
// The operands must have the same type. struct values are compared by their memory address.

// Boolean logical operators require boolean operands.

func (binop *BinOpExpr) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string {
	
	errors = binop.left.TypeCheck(errors, tables, funCallee)
	errors = binop.right.TypeCheck(errors, tables, funCallee)

	lType := binop.left.GetType()
	rType := binop.right.GetType()

	if lType == types.UndeclTySig || rType == types.UndeclTySig {
		binop.ty = types.UnknownTySig
		return errors
	}

	// `>, <, >=, <= `  ==> only integers, return boolean
	if IsRelational(binop.operator){
		if lType == types.IntTySig && rType == types.IntTySig {
			binop.ty = types.BoolTySig
		} else {
			errors = append(errors, SemErr(binop.token,
						fmt.Sprintf("`%s` can only be applied to integer operands.", OpToStr(binop.operator))))
			binop.ty = types.UnknownTySig
		}

	// `+, -, *, / ` ==>  only integers, return integer
	} else if IsArithmetic(binop.operator){
		if lType == types.IntTySig && rType == types.IntTySig {
			binop.ty = types.IntTySig
		} else {
			errors = append(errors, SemErr(binop.token, 
						fmt.Sprintf("`%s` can only be applied to integer operands.", OpToStr(binop.operator))))
			binop.ty = types.UnknownTySig
		}
	
	//` &&, ||`  ==>  only booleans, return boolean
	} else if IsLogical(binop.operator){
		if lType == types.BoolTySig && rType == types.BoolTySig {
			binop.ty = types.BoolTySig
		} else {
			errors = append(errors, SemErr(binop.token,
						fmt.Sprintf("`%s` can only be applied to boolean operands.", OpToStr(binop.operator))))
			binop.ty = types.UnknownTySig
		}

	// `==, !=`   ==> same type, return boolean
	} else if binop.operator == EQ || binop.operator == NEQ {
		if lType == rType {
			// same type but unknown
			if lType == types.UnknownTySig || rType == types.UnknownTySig {
				errors = append(errors, SemErr(binop.token, 
					fmt.Sprintf("`%s` applied to operands of unknown type.", OpToStr(binop.operator))))

				binop.ty = types.UnknownTySig
			} else {
				binop.ty = types.BoolTySig
			}
		// different types
		} else {
			// allow for comparison of struct to nil
			if types.IsStructType(lType) && rType == types.NilTySig || lType == types.NilTySig && types.IsStructType(rType) {
				binop.ty = types.BoolTySig
			} else {
				errors = append(errors, SemErr(binop.token, 
					fmt.Sprintf("`%s` applied to operands of different types: `%s` and `%s`.",
							OpToStr(binop.operator), lType.String(), rType.String())))
									
				binop.ty = types.UnknownTySig
			}
		}
	// operator not recognized; this should never happen
	} else {
		panic("Error in binops typecheck: unrecognized operator.")
	}
	return errors
}


func (binop *BinOpExpr) GetType() types.Type {
	return binop.ty
}


func (binop *BinOpExpr) GetToken() *token.Token{
	return binop.token
}



func (binop *BinOpExpr) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) ir.Operand {
	
	var lOperand, rOperand ir.Operand

	// `nil` operands. Eg: myStruct == nil
	if binop.left.GetType() == types.NilTySig && binop.right.GetType() != types.NilTySig {
		// load struct `nil` registers
		strucEntry, _ := tables.Structs.Contains(binop.right.GetType().String())
		tempReg := ir.NewTempRegister(strucEntry.Ty, funCallee)
		block.AddInstruction(ir.NewLoad(tempReg.GetLabel(), strucEntry.NilRegister.(ir.Operand)))
		lOperand = tempReg

		// load right operand
		rOperand = binop.right.ToLLVM(tables, funCallee, block)

	// `nil` operands. Eg: nil == myStruct
	} else if binop.left.GetType() != types.NilTySig && binop.right.GetType() == types.NilTySig {
		// load left operand
		lOperand = binop.left.ToLLVM(tables, funCallee, block)

		// load struct `nil` registers
		strucEntry, _ := tables.Structs.Contains(binop.left.GetType().String())
		tempReg := ir.NewTempRegister(strucEntry.Ty, funCallee)
		block.AddInstruction(ir.NewLoad(tempReg.GetLabel(), strucEntry.NilRegister.(ir.Operand)))
		rOperand = tempReg
		
	// `nil` operands: nil == nil
	} else if binop.left.GetType() == types.NilTySig && binop.right.GetType() == types.NilTySig {
		// load true
		tempReg := ir.NewTempRegister(types.BoolTySig, funCallee)
		block.AddInstruction(ir.NewLoad(tempReg.GetLabel(), ir.NewLiteral(types.BoolTySig, 1)))
		return tempReg
	
	} else {

		lOperand = binop.left.ToLLVM(tables, funCallee, block)
		rOperand = binop.right.ToLLVM(tables, funCallee, block)
	}

	if IsArithmetic(binop.operator){
		operand := ir.NewTempRegister(binop.ty, funCallee)
		switch binop.operator {
		case ADD:
			block.AddInstruction(ir.NewAdd(operand.GetLabel(), lOperand, rOperand))
		case SUB:
			block.AddInstruction(ir.NewSub(operand.GetLabel(), lOperand, rOperand))
		case MUL:
			block.AddInstruction(ir.NewMul(operand.GetLabel(), lOperand, rOperand))
		case DIV:
			block.AddInstruction(ir.NewSdiv(operand.GetLabel(), lOperand, rOperand))
		}
		return operand
	
	// `>, <, >=, <=, !=, ==` 
	} else if IsRelational(binop.operator) || binop.operator == EQ || binop.operator == NEQ{
		operand := ir.NewTempRegister(binop.ty, funCallee)
		block.AddInstruction(ir.NewIcmp(operand.GetLabel(), OpToLLVM(binop.operator), lOperand, rOperand))
		return operand
	
	//  &&, ||
	} else if IsLogical(binop.operator) {
		operand := ir.NewTempRegister(binop.ty, funCallee)
		
		switch binop.operator {
		case AND:
			block.AddInstruction(ir.NewAnd(operand.GetLabel(), lOperand, rOperand))
		case OR:
			block.AddInstruction(ir.NewOr(operand.GetLabel(), lOperand, rOperand))
			
		default:
			panic("Error in binops to llvm: unrecognized operator.")
		}

		// truncate the boolean result from "i64" to "i1"
		tReg := ir.NewTempRegister(nil, funCallee)
		block.AddInstruction(ir.NewTrunc(tReg.GetLabel(), operand.LLVMTy(), operand, "i1"))
		return tReg

	} else {
		panic("Error in binops to llvm: unrecognized operator.")
	}
}

