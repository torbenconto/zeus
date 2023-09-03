package zeus

import (
	"fmt"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

func LetterCount(in string) int {
	lexer := zeus_lexer.NewLexer(in)
	letterCount := 0

	for {
		token := lexer.NextToken()
		fmt.Println(token)
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if token.Type == zeus_lexer.Word {
			letterCount += len([]rune(token.Value))
		}
	}

	return letterCount
}
