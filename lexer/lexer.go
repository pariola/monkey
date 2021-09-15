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

func (l Lexer) readIdent() token.Token {

	var buf bytes.Buffer

	for {
		r := l.read()

		if !unicode.IsLetter(r) {
			l.unread()
			break
		}

		buf.WriteRune(r)
	}

	return token.Normalize(buf.String())
}

func (l Lexer) readNumber() token.Token {

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

	return token.New(token.NUMBER, buf.String())
}

func (l *Lexer) NextToken() token.Token {

	// skip whitespaces early
	l.skipWhitespace()

	c := l.read()

	switch c {
	case eof:
		return token.New(token.EOF, "")
	case '>':
		return token.New(token.GT, ">")
	case '<':
		return token.New(token.LT, "<")
	case '!':

		if l.peek() == '=' {
			l.read()
			return token.New(token.NOT_EQ, "!=")
		}

		return token.New(token.BANG, "!")
	case '+':
		return token.New(token.PLUS, "+")
	case '-':
		return token.New(token.MINUS, "-")
	case '/':
		return token.New(token.SLASH, "/")
	case ',':
		return token.New(token.COMMA, ",")
	case '(':
		return token.New(token.LPAREN, "(")
	case ')':
		return token.New(token.RPAREN, ")")
	case '{':
		return token.New(token.LBRACE, "{")
	case '}':
		return token.New(token.RBRACE, "}")
	case '=':

		if l.peek() == '=' {
			l.read()
			return token.New(token.EQ, "==")
		}

		return token.New(token.ASSIGN, "=")
	case '*':
		return token.New(token.ASTERISK, "*")
	case ';':
		return token.New(token.SEMICOLON, ";")
	default:
		// handle identifiers & keywords here

		if unicode.IsLetter(c) {
			l.unread()
			return l.readIdent()
		} else if unicode.IsDigit(c) {
			l.unread()
			return l.readNumber()
		}

		return token.New(token.ILLEGAL, "ILLEGAL")
	}
}
