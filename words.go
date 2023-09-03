package zeus

import zeus_lexer "github.com/torbenconto/zeus/lexer"

func WordCount(in string) int {
	lexer := zeus_lexer.NewLexer(in)
	words := 0
	inQuotedText := false

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if token.Type == zeus_lexer.Quotation {
			inQuotedText = !inQuotedText
		}

		if token.Type == zeus_lexer.Word && !inQuotedText {
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
