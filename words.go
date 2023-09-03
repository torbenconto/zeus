package zeus

import zeus_lexer "github.com/torbenconto/zeus/lexer"

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
