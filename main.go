package main

import (
	"fmt"
	"strings"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

func CountSentences(input string) int {
	lexer := zeus_lexer.NewLexer(input)
	sentenceCount := 0

	for {
		token := lexer.NextToken()
		fmt.Println(token)
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if strings.ContainsAny(token.Value, ".!?") {
			sentenceCount++
		}
	}

	return sentenceCount
}

func main() {
	input := `
	This is a test sentence. 
	It contains multiple sentences! Does it really work? 
	Yes, it does. "But how?" you might ask. 
	Well, it handles quotes and abbreviations like Mr. and Dr. as well. 
	And it doesn't get confused by ellipses... or multiple exclamation marks!!!
  `
	sentenceCount := CountSentences(input)
	fmt.Printf("Number of sentences: %d\n", sentenceCount)
}
