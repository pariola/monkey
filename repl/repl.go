package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"monkey/lexer"
	"monkey/token"
)

var (
	prompt = []byte(">> ")
)

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {

		_, _ = out.Write(prompt)

		if !scanner.Scan() {
			return
		}

		line := scanner.Text()

		l := lexer.NewLexer(
			strings.NewReader(line),
		)

		for {

			tok, lit := l.NextToken()

			if tok == token.EOF {
				break
			}

			fmt.Println(tok, lit)
		}
	}
}
