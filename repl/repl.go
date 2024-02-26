package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/w3bdev1/bodmas/lexer"
	"github.com/w3bdev1/bodmas/token"
)

const PROMPT = ">> "

func Start(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	for {
		fmt.Printf("%s", PROMPT)
		ok := scanner.Scan()

		if ok {
			line := scanner.Text()
			lex := lexer.NewLexer(line)
			for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
				fmt.Printf("%+v\n", tok)
			}
		} else {
			err := scanner.Err()
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		}
	}
}
