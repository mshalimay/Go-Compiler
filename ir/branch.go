package ir

import (
	"fmt"
	st "golite/symboltable"
	"golite/codegen"
)

// br i1 <cond>, label <iftrue>, label <iffalse>
// br label <dest>
type Branch struct {
	cond 		string
	labTrue 	int
	labFalse 	int
}


func NewBranch(cond string, labTrue int, labFalse int) *Branch {
	return &Branch{cond, labTrue, labFalse}
}


func (br *Branch) String() string {
	
	// conditional branch
	if br.labFalse != -1 {
		return fmt.Sprintf("br i1 %s, label %%L%d, label %%L%d\n", br.cond, br.labTrue, br.labFalse)
	
	// unconditional branch
	} else {
		return fmt.Sprintf("br label %%L%d\n", br.labTrue)
	}
}

func (br *Branch) Branch() []int {
	return nil
}


func (br *Branch) GetTargets() []int {
	return nil
}


func (br *Branch) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	instrs := codegen.Instructions{}

	if br.cond == "" {
		instrs.WriteComment(br.String())
		instrs.WriteCode(fmt.Sprintf("b .L%v", br.labTrue))
		return instrs.String()
	
	} else {
		instrs.WritePartial(fmt.Sprintf(".L%d\n", br.labTrue))
		instrs.WriteCode(fmt.Sprintf("b .L%d\n", br.labFalse))
		// TODO deal with if - else!
	}
	return instrs.String()
}
