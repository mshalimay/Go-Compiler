package token

import "strconv"

type Token struct {
	Line   int // The line where this token was found
	Column int // The column where this token was found
}

func NewToken(line, column int) *Token {
	return &Token{line, column}
}

func (token *Token) String() string{
	return "(" + strconv.Itoa(token.Line) + ", " + strconv.Itoa(token.Column) + ")"
}
