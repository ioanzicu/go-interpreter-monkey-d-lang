package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ioanzicu/monkeyd/lexer"
	"github.com/ioanzicu/monkeyd/token"
)

// R - READ
// E - EVALUATE
// P - PRINT
// L - LOOP

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// Read input from Terminal
	for {
		fmt.Printf("%v", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		// Pass input to the lexer
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}
