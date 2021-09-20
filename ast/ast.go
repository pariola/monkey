package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token // IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatement struct {
	Token token.Token // LET
	Name  *Identifier
	Value Expression
}

func (s *LetStatement) statementNode() {}

func (s *LetStatement) TokenLiteral() string { return s.Token.Literal }

type ReturnStatement struct {
	Token token.Token // RETURN
	Name  *Identifier
	Value Expression
}

func (s *ReturnStatement) statementNode() {}

func (s *ReturnStatement) TokenLiteral() string { return s.Token.Literal }
