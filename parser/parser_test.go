package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"monkey/ast"
	"monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.NewLexer(
		strings.NewReader(input),
	)

	p := NewParser(l)
	program := p.Parse()

	if program == nil {
		t.Fatalf("Parse() returned nil")
	}

	require.Equal(t, 3, len(program.Statements), "program.Statements does not contain 3 statements")

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10.3;
return add(1, 4);
`
	l := lexer.NewLexer(
		strings.NewReader(input),
	)

	p := NewParser(l)
	program := p.Parse()

	if program == nil {
		t.Fatalf("Parse() returned nil")
	}

	require.Equal(t, 3, len(program.Statements), "program.Statements does not contain 3 statements")

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		assert.True(t, ok, "statement not a ReturnStatement")

		assert.Equal(t, "return", returnStmt.TokenLiteral())
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {

	if assert.Equal(t, "let", s.TokenLiteral()) {
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if assert.True(t, ok, "statement not a LetStatement") {
		return false
	}

	if assert.Equal(t, letStmt.Name.Value, name) || assert.Equal(t, letStmt.Name.TokenLiteral(), name) {
		return false
	}

	return true
}
