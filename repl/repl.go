// Package repl provides a Read Eval Print Loop for the Monkey lanaguage.
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pto/monkey/lexer"
	"github.com/pto/monkey/token"
)

// PROMPT is the REPL command prompt.
const PROMPT = ">> "

// Start begins a REPL from in to out.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
