package ast

import(
	"bytes"
	"golite/token"
)

// NOTE: this node may be deleted from the AST, but helps in modularity/simplicity in parser code

type Parameters struct{
	token		    *token.Token
	params 			[]*Decl
}

func NewParameters(token *token.Token, params []*Decl) *Parameters{
	return &Parameters{token, params}
}

func (p *Parameters) String() string{
	// parameters = '(' [ decl {',' decl}] ')' 
	// decl: ID type // eg: a int // b bool

	var out bytes.Buffer

	out.WriteString("(")

	// iterate through decls and add them to the string
	count := 0
	for _, decl := range p.params{
		// write eg: "a int"
		out.WriteString(decl.id)
		out.WriteString(" ")
		out.WriteString(decl.id)
		// if not last declaration, add a COMMA before next declaration
		if count != len(p.params) - 1{
			out.WriteString(", ")
		}
		count++
	}
	out.WriteString(")")

	return out.String()  // eg: (a int, b bool, c *int)
}
