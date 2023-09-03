package zeus

import zeus_lexer "github.com/torbenconto/zeus/lexer"

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

func GetAllWords(input string) []string {
	lexer := zeus_lexer.NewLexer(input)
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
