package ast

// inherits from Statement

import (
	"bytes"
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Ret struct {
	token 			*token.Token
	expression 		Expression
	ty 				types.Type
}

func NewRet(token *token.Token, expression Expression, ty types.Type) *Ret{
	return &Ret{token, expression, ty}
}

func (ret *Ret) String() string{
	// Return: 'return' [Expression] ';'
	
	var out bytes.Buffer

	out.WriteString("return")

	if ret.expression != nil {
		out.WriteString(" ")
		out.WriteString(ret.expression.String())
	}
	
	out.WriteString(";\n")
	
	return out.String()
}


func(ret *Ret) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{

	// asssign type for the `ret` statement
	if ret.expression != nil {
		errors = ret.expression.TypeCheck(errors, tables, funCallee)
		ret.ty = ret.expression.GetType()
	} else {
		ret.ty = types.VoidTySig
	}

	// check if the return type matches the callee function's return type
	expecRetTy := funCallee.RetTy

	if expecRetTy == types.VoidTySig && ret.ty != types.VoidTySig {
			errors = append(errors, SemErr(ret.token, 
					fmt.Sprintf("too many return values in return for `%s`. Got: `%s`.",
					funCallee.Name, ret.ty.String())))

	
	} else if expecRetTy != types.VoidTySig && ret.ty == types.VoidTySig {
			errors = append(errors, SemErr(ret.token, 
							fmt.Sprintf("not enough return values in return for `%s`. Expected: `%s`.",
									funCallee.Name, expecRetTy.String())))
	
	} else if expecRetTy != ret.ty {
			errors = append(errors, SemErr(ret.token, 
							fmt.Sprintf("type mismatch in return for `%s`. Expected: `%s`. Got: `%s`.",
									funCallee.Name, expecRetTy.String(), ret.ty.String())))
	}
		
	return errors
}

func (ret *Ret) GetType() types.Type{
	return ret.ty
}

func (ret *Ret) HasReturn(warnings []string) ([]string, bool){
	return warnings, true
}

func (ret *Ret) GetToken() *token.Token{
	return ret.token
}


func (ret *Ret) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) *ir.BasicBlock {
	// return newPt;	
	// %r37 = load %struct.Point2D*, %struct.Point2D** %newpt
    // store %struct.Point2D* %r37, %struct.Point2D** %_retval
    // %r38 = load %struct.Point2D*, %struct.Point2D** %_retval
    // ret %struct.Point2D* %r38

	if ret.expression != nil {
		block.AddInstruction(ir.NewComment("return "  + ret.expression.String()))
	} else {
		block.AddInstruction(ir.NewComment("return"))
	}

	if ret.expression == nil && funCallee.Name != "main" {
		block.AddInstruction(ir.NewRet(nil))
	
	} else if funCallee.Name == "main" {
		// store i64 0, i64* %_mainretval
		block.AddInstruction(ir.NewStore(ir.NewLiteral(types.IntTySig, 0), funCallee.RetValReg.(ir.Operand)))
		
		// %r41 = load i64, i64* %_mainretval
		reg := ir.NewTempRegister(types.IntTySig, funCallee)
		block.AddInstruction(ir.NewLoad(reg.GetLabel(), funCallee.RetValReg.(ir.Operand)))

		// ret i64 %r41
		block.AddInstruction(ir.NewRet(reg))
	} else {
		// load result from `expression` into register
		reg := ret.expression.ToLLVM(tables, funCallee, block)

		// store result into function's retval register
		block.AddInstruction(ir.NewStore(reg, funCallee.RetValReg.(ir.Operand)))

		// load result from function's retval register into register
		reg2 := ir.NewTempRegister(ret.ty, funCallee)
		block.AddInstruction(ir.NewLoad(reg2.GetLabel(), funCallee.RetValReg.(ir.Operand)))

		// add return instruction
		block.AddInstruction(ir.NewRet(reg2))
	}
	
	return block
}
