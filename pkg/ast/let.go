package ast

import "github.com/estevesnp/dsb/pkg/token"

type LetStatement struct {
	Name  *Identifier
	Value Expression
	Token token.Token // token.LET
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
