package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ioanzicu/monkeyd/evaluator"
	"github.com/ioanzicu/monkeyd/lexer"
	"github.com/ioanzicu/monkeyd/parser"
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
		if line == "exit" {
			fmt.Println("Bye, bye ...!")
			return
		}

		// Pass input to the lexer
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MONKEY_FACE = `
      .--.  .-"__,__"-. .--.
     / .. \/  .-. .-.  \/ .. \
    | |  '|  /   Y   \  |'  | |
    | \   \  \ 0 | 0 /  /   / |
     \ '-, \.-"""""""-./,-'  /
      ''-' /_   ^ ^   _\ '-''
          | \._    _./  |
          \  \  '~' /   /
           '._ '-=-' _.'
              '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
