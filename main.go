package main

import (
	"fmt"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

func main() {
	input := `Hello, World! This is a "simple text" lexer.`

	lexer := zeus_lexer.NewLexer(input)

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}
		fmt.Printf("Type: %v, Value: %s\n", token.Type, token.Value)
	}
}
