package token

type Token struct {
	Type    Type
	Literal string
}

func New(t Type, l string) Token {
	return Token{Type: t, Literal: l}
}

type Type int

const (
	ILLEGAL Type = iota
	EOF

	// Identifiers, Numbers & literals
	NUMBER // 1, 12443, 12.325
	IDENT  // add, foobar

	// Operators
	PLUS     // +
	BANG     // !
	MINUS    // -
	ASSIGN   // =
	SLASH    // /
	ASTERISK // *

	// Comparision
	GT     // >
	LT     // <
	EQ     // ==
	NOT_EQ // !=

	// Delimiters
	COMMA     // ,
	SEMICOLON // ;

	// Braces & Parenthesis
	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }

	// Keywords
	LET      // let
	FUNCTION // fn
	TRUE     // true
	FALSE    // false
	IF       // if
	ELSE     // else
	RETURN   // return
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func Normalize(ident string) Token {

	if t, ok := keywords[ident]; ok {
		return New(t, ident)
	}

	return New(IDENT, ident)
}
