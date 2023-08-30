package ast

// inherits from Statement

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)


type Loop struct {
	token 			*token.Token
	expression 		Expression
	block 			Statement
}


func NewLoop(token *token.Token, expression Expression, block Statement) *Loop{
	return &Loop{token, expression, block}
}


func (loop *Loop) String() string{
	// Loop: 'for' '(' Expression ')' Block 
	var out bytes.Buffer

	out.WriteString("for ")
	out.WriteString("(")
	out.WriteString(loop.expression.String())
	out.WriteString(")")
	out.WriteString(loop.block.String())

	return out.String()
}


func (loop *Loop) TypeCheck(errors []string, tables *st.SymbolTables, funCallee *st.FuncEntry) []string{
	// typecheck the boolean expression and get it's type
	errors = loop.expression.TypeCheck(errors, tables, funCallee)
	exprType := loop.expression.GetType()

	// if expression is not of type bool, add error
	if exprType != types.BoolTySig {
		errors = append(errors, SemErr(loop.token, "expression in `for` must be of type `bool`. Got: "+
										exprType.String() + "."))
	}

	errors = loop.block.TypeCheck(errors, tables, funCallee)
	return errors
}

func (loop *Loop) HasReturn(warnings []string) ([]string, bool){
	// although might have returns inside loop, not guaranteed
	// to enter loop at compile time, so always false.
	// get the warnings from the block though (unreachable code)

	warnings, _ = loop.block.HasReturn(warnings)
	return warnings, false
}

func (loop *Loop) GetToken() *token.Token{
	return loop.token
}


func (loop *Loop) ToLLVM(tables *st.SymbolTables, funCallee *st.FuncEntry, block *ir.BasicBlock) *ir.BasicBlock{
	
	//-----------------------------------------------------------------------
	// Create and link entry block, body block
	//-----------------------------------------------------------------------
	// create ENTRY block for `for` statement
	entryBlock := ir.NewBasicBlock(ir.FOR_ENTRY)
	entryBlock.AddInstruction(ir.NewComment("for "+ loop.expression.String()))
 
	// `entryBlock` is successor of `block`; `block` is predecessor of `entryBlock`
	block.AddSuccessor(entryBlock)
	entryBlock.AddPredecessor(block)

	// create `body` block; `body` block is 1st successor of `entryBlock`; `entryBlock` is predecessor of `trueBlock`
	bodyBlock := ir.NewBasicBlock(ir.FOR_BODY)
	entryBlock.AddSuccessor(bodyBlock)
	bodyBlock.AddPredecessor(entryBlock)

	// ----------------------------------------------------------------------
	// Fill "caller" block
	//-----------------------------------------------------------------------
	// add unconditional branching instruction to the caller `block`
	// eg: br label %L3
	block.AddInstruction(ir.NewBranch("", entryBlock.GetLabel(), -1))

	// -----------------------------------------------------------------------
	// get result from conditional expression to fill entry block at the end
	// -----------------------------------------------------------------------
	// obs1: we get the result here to preserve the order of labels of virtual registers
	// obs2: we fill the entry block at the end to preserve the order of of labels for blocks
	// above is because in the `body` block, new virtual register/blocks may be created
	// notice: diff from IF stmt, because there the icmp is to the true/else blocks created previously, not the `exit` block
	reg := Expression.ToLLVM(loop.expression, tables, funCallee, entryBlock)

	
	// obs: the branch below deals with the case "for (myBooleanFun())  {...}"
	// that is, when there is only a single term and not a boolean expression
	// a bitcast to "i1" is necessary because in this case, the function that
	// returns a boolean and we are treating booleans as "i64" here. 
	// In other cases, the ICMP returns "i1" already so no need for bitcast.
	// Ex: not need a bitcast: for (myBooleanFun() && false); for (a>1); ...
	lastIns := entryBlock.GetLastInstruction()
	if _, ok := lastIns.(*ir.Icmp); !ok {
		tReg := ir.NewTempRegister(nil , funCallee)
		entryBlock.AddInstruction(ir.NewTrunc(tReg.GetLabel(), reg.LLVMTy(), reg, "i1"))
		reg = tReg
	}
	// -----------------------------------------------------------------------
	// Fill body block
	// -----------------------------------------------------------------------
	finalBodyBlock := bodyBlock
	finalBodyBlock = loop.block.ToLLVM(tables, funCallee, finalBodyBlock)
	
	// fill body block with unconditional branch instruction
	finalBodyBlock.AddInstruction(ir.NewBranch("", entryBlock.GetLabel(), -1))

	// -----------------------------------------------------------------------
	// Create exit block; fill entry block with conditional branch instruction
	// -----------------------------------------------------------------------
	exitBlock := ir.NewBasicBlock(ir.FOR_EXIT)

	// `exitBlock` is 2nd successor of `entryBlock`; `entryBlock` is predecessor of `exitBlock`
	entryBlock.AddSuccessor(exitBlock)

	// fill entry block with conditional branch instruction
	entryBlock.AddInstruction(ir.NewBranch(reg.String(), bodyBlock.GetLabel(), exitBlock.GetLabel()))

	return exitBlock
}

