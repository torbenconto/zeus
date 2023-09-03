package zeus

import (
	"strings"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

// SentenceCount counts the total number of sentences in a text using the lexer,
// considering abbreviations.
func SentenceCount(in string) int {
	lexer := zeus_lexer.NewLexer(in)
	sentenceCount := 0
	inQuotedText := false
	lastToken := zeus_lexer.Token{Type: zeus_lexer.Unknown, Value: ""}

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if token.Type == zeus_lexer.Quotation {
			inQuotedText = !inQuotedText
		}

		if !inQuotedText && (strings.ContainsAny(token.Value, ".!?") && token.Type != zeus_lexer.Abbreviation) {
			if lastToken.Type != zeus_lexer.Punctuation && lastToken.Type != zeus_lexer.Comma {
				sentenceCount++
			}
		}

		lastToken = token
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
