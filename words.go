package zeus

import (
	"fmt"
	"strings"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

// RemoveExtraWhitespace removes extra whitespace between words using the lexer.
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

	fmt.Println(result.String())

	return result.String()
}

// Returns the amount of words in a given string (in)
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

// Returns a list of all words in a given string
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
