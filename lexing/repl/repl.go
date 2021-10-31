package repl

import (
	`bufio`
	`fmt`
	`io`

	`lexing/lexer`
	`lexing/token`
)

const PROMPT = ">> "

func Start(in io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%v\n", tok)
		}
	}
}
