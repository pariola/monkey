package lexer

import (
	"bufio"
	"bytes"
	"io"
	"unicode"

	"monkey/token"
)

const (
	eof = rune(0)
)

type Lexer struct {

	// position tracker
	pos int32

	// reader
	r *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		r: bufio.NewReader(r),
	}
}

func (l *Lexer) peek() rune {

	r := l.read()
	defer l.unread()

	return r
}

func (l *Lexer) read() rune {

	l.pos++

	r, _, err := l.r.ReadRune()

	if err != nil {
		l.pos = 0
		return eof
	}

	return r
}

func (l *Lexer) unread() {

	if l.pos > 0 {
		l.pos--
		_ = l.r.UnreadRune()
	}
}

func (l Lexer) skipWhitespace() {

	for {
		r := l.read()

		if !unicode.IsSpace(r) {
			l.unread()
			break
		}
	}
}

func (l Lexer) readIdent() (token.Token, string) {

	var buf bytes.Buffer

	for {
		r := l.read()

		if !unicode.IsLetter(r) {
			l.unread()
			break
		}

		buf.WriteRune(r)
	}

	ident := buf.String()
	return token.Normalize(ident), ident
}

func (l Lexer) readNumber() (token.Token, string) {

	var hasPoint bool
	var buf bytes.Buffer

	for {
		r := l.read()

		if r == '.' && !hasPoint {
			hasPoint = true
		} else if !unicode.IsDigit(r) {
			l.unread()
			break
		}

		buf.WriteRune(r)
	}

	return token.NUMBER, buf.String()
}

func (l *Lexer) NextToken() (token.Token, string) {

	// skip whitespaces early
	l.skipWhitespace()

	c := l.read()

	switch c {
	case eof:
		return token.EOF, ""
	case '>':
		return token.GT, ">"
	case '<':
		return token.LT, "<"
	case '!':

		if l.peek() == '=' {
			l.read()
			return token.NOT_EQ, "!="
		}

		return token.BANG, "!"
	case '+':
		return token.PLUS, "+"
	case '-':
		return token.MINUS, "-"
	case '/':
		return token.SLASH, "/"
	case ',':
		return token.COMMA, ","
	case '(':
		return token.LPAREN, "("
	case ')':
		return token.RPAREN, ")"
	case '{':
		return token.LBRACE, "{"
	case '}':
		return token.RBRACE, "}"
	case '=':

		if l.peek() == '=' {
			l.read()
			return token.EQ, "=="
		}

		return token.ASSIGN, "="
	case '*':
		return token.ASTERISK, "*"
	case ';':
		return token.SEMICOLON, ";"
	default:
		// handle identifiers & keywords here

		if unicode.IsLetter(c) {
			l.unread()
			return l.readIdent()
		} else if unicode.IsDigit(c) {
			l.unread()
			return l.readNumber()
		}

		return token.ILLEGAL, "ILLEGAL"
	}
}
