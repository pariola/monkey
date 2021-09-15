package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// load 2 tokens
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() *ast.Program {

	program := &ast.Program{
		Statements: make([]ast.Statement, 0),
	}

	for p.curToken.Type != token.EOF {
		if stmt := p.parseStatement(); stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {

	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {

	s := &ast.LetStatement{Token: p.curToken}

	if p.peekToken.Type != token.IDENT {
		return nil
	}

	p.nextToken()

	s.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if p.peekToken.Type != token.ASSIGN {
		return nil
	}

	p.nextToken()

	// todo: read value

	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return s
}
