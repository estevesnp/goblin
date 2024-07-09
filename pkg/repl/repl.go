package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/estevesnp/dsb/pkg/lexer"
	"github.com/estevesnp/dsb/pkg/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scan := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		if scanned := scan.Scan(); !scanned {
			return
		}

		line := scan.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
