package ast

import "golite/token"

type Decl struct {
	token *token.Token
	id    string
	ty    string
}

func NewDecl(token *token.Token, id string, ty string) *Decl {
	return &Decl{token, id, ty}
}

func (decl *Decl) GetToken() *token.Token {
	return decl.token
}
