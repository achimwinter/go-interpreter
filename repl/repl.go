package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/achimwinter/monkey-language/evaluator"
	"github.com/achimwinter/monkey-language/lexer"
	"github.com/achimwinter/monkey-language/object"
	"github.com/achimwinter/monkey-language/parser"
)

const SNOWMAN = `
 *   *   *
* * * * * *
 *   *   *
   
     _  
    (.) 
  <( : )>
   ( : )
`

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned:= scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, SNOWMAN)
	io.WriteString(out, "Woops, things getting icy. We ran into an error!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}
