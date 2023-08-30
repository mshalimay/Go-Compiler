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


type Assignment struct {
	token      *token.Token
	lvalue     Expression
	expression Expression
	scan       *Scan
	ty  		types.Type
}

func NewAssignment(token *token.Token, lvalue Expression, expression Expression, scan *Scan) *Assignment {
	return &Assignment{token, lvalue, expression, scan, nil}
}

func (assignment *Assignment) String() string {
	// Assignment: LValue '=' ( Expression | 'scan' ) ';'
	// LValue = 'id' {'.' id}

	var out bytes.Buffer

	out.WriteString(assignment.lvalue.String())

	out.WriteString(" = ")

	if assignment.expression != nil {
		out.WriteString(assignment.expression.String())
	} else if assignment.scan != nil {
		out.WriteString(assignment.scan.String())
	}

	out.WriteString(";\n")
	return out.String()
}

func (a *Assignment) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string {
	// assignment: lvalue EQUALS (expression | SCAN) SEMICOLON;
	// lvalue: ID lvaluePrime*;
	// a = 1; a = scan; a.b = x.y.w.z * 3

	// typecheck lvale and get it's type
	errors = a.lvalue.TypeCheck(errors, tables, funCallee)
	lvalueType := a.lvalue.GetType()
	a.ty = lvalueType

	if lvalueType == types.UndeclTySig {
		errors = append(errors, SemErr(a.token, "LHS of assignment not declared"))
	}

	// case 1: a = scan; a.b.c = scan
	if a.expression == nil {
		if lvalueType != types.IntTySig {
			//NOTE: lvalueType can be `nil` if the variable in the LHS is not declared
			errors = append(errors, SemErr(a.token, "cannot assign `scan` to non-integer type: `"+lvalueType.String()+"`."))
		}
		return errors
	}

	// case 2: a = 1; a.b.c = x.y.w.z * 3

	// typecheck rhs expression and get it's type
	errors = a.expression.TypeCheck(errors, tables, funCallee)
	rvalueType := a.expression.GetType()

	if rvalueType == types.UndeclTySig {
		errors = append(errors, SemErr(a.token, "RHS of assignment not declared"))
	}

	// NOTE: dont need to return assignment type error if undeclared types
	if lvalueType == types.UndeclTySig || rvalueType == types.UndeclTySig {
		return errors
	}

	if lvalueType != rvalueType {
		if types.IsStructType(lvalueType) && rvalueType == types.NilTySig {
		} else {
			errors = append(errors, SemErr(a.token,
				fmt.Sprintf("type mismatch in assignment: cannot assign `%s` to `%s`", rvalueType.String(), lvalueType.String())))
		}
	}
	return errors
}

func (a *Assignment) HasReturn(warnings []string) ([]string, bool) {
	return warnings, false
}

func (a *Assignment) GetToken() *token.Token {
	return a.token
}

func (a *Assignment) ToLLVM(tables *st.SymbolTables, funcCallee *st.FuncEntry, block *ir.BasicBlock) *ir.BasicBlock {
	
	block.AddInstruction(ir.NewComment(strings.Replace(a.String(), "\n", "", -1)))	
	
	// case 1: a = scan; a.b.c = scan
	if scan := a.scan; scan != nil {
		destReg := a.lvalue.ToLLVM(tables, funcCallee, block)
		block.AddInstruction(ir.NewScanf(destReg))
	
	// case 2: myStruc = nil
	} else if a.expression.GetType() == types.NilTySig {
		// typecast lvalue to struct
		strucName := a.lvalue.GetType().String()
		strucEntry, _ := tables.Structs.Contains(strucName)
		
		newReg := ir.NewTempRegister(strucEntry.Ty, funcCallee)
		block.AddInstruction(ir.NewLoad(newReg.GetLabel(), strucEntry.NilRegister.(ir.Operand)))
		destReg := a.lvalue.ToLLVM(tables, funcCallee, block)
		block.AddInstruction(ir.NewStore(newReg, destReg))
	
	} else {
		rValue := a.expression.ToLLVM(tables, funcCallee, block)
		destReg := a.lvalue.ToLLVM(tables, funcCallee, block)
		block.AddInstruction(ir.NewStore(rValue, destReg))		
	}
	return block	
}