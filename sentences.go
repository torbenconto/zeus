package zeus

import (
	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

// Returns the amount of sentences in the given string (in)
func SentenceCount(in string) int {
	lexer := zeus_lexer.NewLexer(in)
	sentenceCount := 0

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if token.Type == zeus_lexer.Punctuation {
			sentenceCount++
		}
	}

	return sentenceCount
}

// Returns the number of words in the given string (in) over the number of sentences in the given string (in). If the number of sentences in the given string is 0, the number of words is returned.
func AverageWordsPerSentence(in string) float64 {

	sentences := SentenceCount(in)
	words := WordCount(in)

	if sentences == 0 {
		return float64(words)
	}

	return float64(words) / float64(sentences)
}
