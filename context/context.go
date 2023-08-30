package context

import (
	"fmt"
)

type CompilerPhase int

const (
	LEXER CompilerPhase = iota // 0 - LEXER, 1 - PARSER, ....
	PARSER
)


func(err *CompilerError) String() string{
	if err.Phase == LEXER {
		return fmt.Sprintf("Lexer error(%d, %d): %s", err.Line, err.Column, err.Msg)
	} else if err.Phase == PARSER {
		return fmt.Sprintf("Parser error(%d, %d): %s", err.Line, err.Column, err.Msg)
	}
	panic("Invalid phase found")
}

type CompilerError struct {
	Line int
	Column int
	Msg string
	Phase CompilerPhase
}

func HasErrors(errors []*CompilerError) bool{
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		return true
	}
	return false
}