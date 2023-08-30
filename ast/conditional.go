package ast

// inherits from Statement

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type Conditional struct {
	token      	*token.Token
	expression 	Expression
	block      	Statement
	elseBlock  	Statement
}

func NewConditional(token *token.Token, expression Expression, block Statement, elseBlock Statement) *Conditional {
	return &Conditional{token, expression, block, elseBlock}
}

func (conditional *Conditional) String() string {
	// Conditional:'if' '(' Expression ')' Block ['else' Block]

	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString("(")
	out.WriteString(conditional.expression.String())
	out.WriteString(")")
	out.WriteString(conditional.block.String())

	if conditional.elseBlock != nil {
		// remove previous "\n" if any
		index := bytes.LastIndex(out.Bytes(), []byte("\n"))

		if index != -1 {
			// Replace everything after the last "\n" (including it) with " else "
			out.Truncate(index)
			out.WriteString(" else ")
		} else {
			// add " else " if no "\n"
			out.WriteString(" else ")
		}
		out.WriteString(conditional.elseBlock.String())
	}

	return out.String()
}

func (cond *Conditional) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string {

	// typecheck the boolean expression and get it's type
	errors = cond.expression.TypeCheck(errors, tables, funCallee)
	exprType := cond.expression.GetType()

	// if expression is not of type bool, add error
	if exprType != types.BoolTySig {
		errors = append(errors, SemErr(cond.token, "expression in `if` must be of type bool. Got: "+
			exprType.String()+"."))
	}

	// typecheck conditional block and else block if any and add errors
	errors = cond.block.TypeCheck(errors, tables, funCallee)

	if cond.elseBlock != nil {
		errors = cond.elseBlock.TypeCheck(errors, tables, funCallee)
	}

	return errors
}


// HasReturn is `true` if `cond`: (i) has `if` and `else` branches AND both have return.
// That is, returns `true` only if the if-else is "closed" (guranteed to return).
// note: if no `else` branch, cannot guarantee return, since `if` branch may not be taken
func (cond *Conditional) HasReturn(warnings []string) ([]string, bool) {
	hasReturn := true
	var b bool

	// if block inside IF has return, then IF has return
	warnings, b = cond.block.HasReturn(warnings)
	hasReturn = hasReturn && b

	// if else block is present, check if it has return; otherwise, `cond` has no return
	if cond.elseBlock != nil {
		warnings, b = cond.elseBlock.HasReturn(warnings)
		hasReturn = hasReturn && b
	} else {
		hasReturn = false
	}

	return warnings, hasReturn
}

func (cond *Conditional) GetToken() *token.Token {
	return cond.token
}


func (cond *Conditional) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) *ir.BasicBlock {
	
	//-----------------------------------------------------------------------
	// Create and link entry, true, else blocks
	//-----------------------------------------------------------------------
	// create ENTRY block for `if` statement
	entryBlock := ir.NewBasicBlock(ir.IF_ENTRY)
	entryBlock.AddInstruction(ir.NewComment("If " + cond.expression.String()))
 
	// `entryBlock` is successor of `block`; `block` is predecessor of `entryBlock`
	block.AddSuccessor(entryBlock)
	entryBlock.AddPredecessor(block)

	// create TRUE block; trueBlock` is 1st (lhs) successor of `entryBlock`; `entryBlock` is predecessor of `trueBlock`
	trueBlock := ir.NewBasicBlock(ir.IF_TRUE)
	entryBlock.AddSuccessor(trueBlock)
	trueBlock.AddPredecessor(entryBlock)

	// ELSE block; `elseBlock` is 2nd (rhs) successor of `entryBlock`; `entryBlock` is predecessor of `elseBlock`
	elseBlock := ir.NewBasicBlock(ir.IF_ELSE)
	entryBlock.AddSuccessor(elseBlock)
	elseBlock.AddPredecessor(entryBlock)

	// ----------------------------------------------------------------------
	// Fill "caller" block
	//-----------------------------------------------------------------------
	// add branching instruction to the caller `block`
	// eg: br label %L3
	block.AddInstruction(ir.NewBranch("", entryBlock.GetLabel(), -1))

	// -----------------------------------------------------------------------
	// Fill `if` entry block
	// -----------------------------------------------------------------------
	// fill with condition // eg: if (x > 0)

	// obs: the branch below deals with the case "if (myBooleanFun())  {...}"
	// when there is only a single term and not a boolean expression a `trunc` to "i1" 
	// is necessary because the boolean returned by the function is an "i64"
	// in this implementation; but we need "i1" for the branch instruction.
	// In other cases, a ICMP will return "i1" so no need for bitcast.
	// Ex: no need for bitcast: if (myBooleanFun() && false); if (a>1); ...
	reg := Expression.ToLLVM(cond.expression, tables, funCallee, entryBlock)

	lastIns := entryBlock.GetLastInstruction()
	if _, ok := lastIns.(*ir.Icmp); !ok {
		tReg := ir.NewTempRegister(nil , funCallee)
		entryBlock.AddInstruction(ir.NewTrunc(tReg.GetLabel(), reg.LLVMTy(), reg, "i1"))
		reg = tReg
	}
	entryBlock.AddInstruction(ir.NewBranch(reg.String(), trueBlock.GetLabel(), elseBlock.GetLabel()))

	// -----------------------------------------------------------------------
	// Fill `true` block
	// -----------------------------------------------------------------------
	// REVIEW why the "finalTrueBlock"?
	trueBlock.AddInstruction(ir.NewComment("True " + cond.expression.String()))
	finalTrueBlock := trueBlock
	finalTrueBlock = cond.block.ToLLVM(tables, funCallee, trueBlock)

	// -----------------------------------------------------------------------
	// Fill `else` block
	// -----------------------------------------------------------------------
	finalElseBlock := elseBlock
	if cond.elseBlock != nil {
		finalElseBlock.AddInstruction(ir.NewComment("Else " +  cond.expression.String()))
		finalElseBlock = cond.elseBlock.ToLLVM(tables, funCallee, elseBlock)
	}

	// -----------------------------------------------------------------------
	// Create exit block and add terminators for true/else blocks, if needed
	// -----------------------------------------------------------------------
	var trueHasRet, elseHasRet bool = false, false

	trueHasRet = (cond.block.(*Block)).HasRet
	if cond.elseBlock != nil {
		elseHasRet = (cond.elseBlock.(*Block)).HasRet
	}

	if trueHasRet && elseHasRet {
		// if both true and else blocks have return, no need for exit block
		return block
	
	} else {
		exitBlock := ir.NewBasicBlock(ir.IF_EXIT)
		exitBlock.AddInstruction(ir.NewComment("Exit " + cond.expression.String()))

		// `exitBlock` is 3rd successor of `entryBlock`; `entryBlock` is predecessor of `exitBlock`
		entryBlock.AddSuccessor(exitBlock)
		exitBlock.AddPredecessor(entryBlock)

		// if no terminator in true branch, add unconditional branch to exit block
		if !trueHasRet {
			// if no return statement in block, add unconditional branch to exit block
			finalTrueBlock.AddInstruction(ir.NewBranch("", exitBlock.GetLabel(), -1))
		}

		// if no terminator in else branch, add unconditional branch to exit block
		if !elseHasRet {
			finalElseBlock.AddInstruction(ir.NewBranch("", exitBlock.GetLabel(), -1))
		}

		return exitBlock
	}
}


