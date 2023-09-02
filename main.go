package main

import (
	"fmt"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

func main() {
	input := `Hello, World! This isn't a "simple text" lexer.`

	lexer := zeus_lexer.NewLexer(input)
	sentenceCount := 0

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		// Check if the token is a punctuation mark that typically ends a sentence.
		if token.Type == zeus_lexer.Punctuation {
			sentenceCount++
		}

		fmt.Printf("Type: %v, Value: %s\n", token.Type, token.Value)
	}

	fmt.Printf("Total Sentences: %d\n", sentenceCount)
}
