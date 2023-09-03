package zeus

import (
	"fmt"
	"strings"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

func SentenceCount(in string) int {
	lexer := zeus_lexer.NewLexer(in)
	sentenceCount := 0

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if token.Type == zeus_lexer.Punctuation {
			fmt.Println(token)
		}

		if token.Type == zeus_lexer.Punctuation &&
			(strings.ContainsAny(token.Value, ".!?")) {
			sentenceCount++
		}
	}

	return sentenceCount
}

func AverageWordsPerSentence(in string) float64 {

	sentences := SentenceCount(in)
	words := WordCount(in)

	if sentences == 0 {
		return float64(words)
	}

	return float64(words) / float64(sentences)
}
