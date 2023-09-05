package zeus

import (
	"strings"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

// RemoveExtraWhitespace removes extra whitespace between words in the given string (in) using the lexer.
//
// Parameters:
//   - in: The input string from which to remove extra whitespace.
//
// Returns:
//   - The input string with extra whitespace removed.
func RemoveExtraWhitespace(in string) string {
	lexer := zeus_lexer.NewLexer(in)
	var result strings.Builder
	lastToken := zeus_lexer.Token{Type: zeus_lexer.Unknown, Value: ""}

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		// Only add a space if the current token is a Word and the previous token was not punctuation.
		if token.Type == zeus_lexer.Word && lastToken.Type != zeus_lexer.Punctuation {
			result.WriteRune(' ')
		}

		result.WriteString(token.Value)

		lastToken = token
	}

	return result.String()[1:]
}

// WordCount returns the number of words in the given string (in).
//
// Parameters:
//   - in: The input string for which to count words.
//
// Returns:
//   - The number of words in the input string.
func WordCount(in string) int {
	lexer := zeus_lexer.NewLexer(in)
	words := 0

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if token.Type == zeus_lexer.Word {
			words++
		}
	}

	return words
}

// GetAllWords returns a list of all words in the given string.
//
// Parameters:
//   - in: The input string from which to extract words.
//
// Returns:
//   - A slice of strings containing all the words in the input string.
func GetAllWords(in string) []string {
	lexer := zeus_lexer.NewLexer(in)
	words := []string{}

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if token.Type == zeus_lexer.Word {
			words = append(words, token.Value)
		}
	}

	return words
}
