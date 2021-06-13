package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/moll-y/monkey/pkg/lexer"
	"github.com/moll-y/monkey/pkg/token"
)

const prompt = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)
		scanned := scanner.scan()

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
