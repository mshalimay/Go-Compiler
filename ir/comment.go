package ir

import (
	"fmt"
	st "golite/symboltable"
)



type Comment struct {
	Comment string
}

func NewComment(comment string) *Comment {
	return &Comment{Comment: comment}
}


func (c *Comment) String() string {
	return fmt.Sprintf("; %s\n", c.Comment)
}

func (c *Comment) GetTargets() []int {
	return nil
}


func (c *Comment) TranslateToAssembly(tables *st.SymbolTables, funCallee *st.FuncEntry) string {
	return ""
}