package zeus

import (
	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

// LetterCount counts the number of letters (alphabetic characters) in a given input string.
//
// Parameters:
//   - in: The input string for which you want to count the letters.
//
// Returns:
//   - The count of letters in the input string.
func LetterCount(in string) int {
	lexer := zeus_lexer.NewLexer(in)
	letterCount := 0

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if token.Type == zeus_lexer.Word {
			letterCount += len([]rune(token.Value))
		}
	}

	return letterCount
}
