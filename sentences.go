package zeus

import (
	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

// SentenceCount returns the number of sentences in the given string (in).
// Note that this function may count unknown characters and abbreviations as sentences, which may not be accurate.
//
// Parameters:
//   - in: The input string for which to count sentences.
//
// Returns:
//   - The number of sentences in the input string.
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

// AverageWordsPerSentence calculates the average number of words per sentence in the given string (in).
// If the number of sentences in the given string is 0, it returns the number of words as a float64.
//
// Parameters:
//   - in: The input string for which to calculate the average words per sentence.
//
// Returns:
//   - The average words per sentence as a float64.
func AverageWordsPerSentence(in string) float64 {
	sentences := SentenceCount(in)
	words := WordCount(in)

	if sentences == 0 {
		return float64(words)
	}

	return float64(words) / float64(sentences)
}
