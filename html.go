package zeus

import (
	"fmt"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

func GetAllHTMLTags(in string) []string {
	lexer := zeus_lexer.NewLexer(in)
	var htmlTags []string

	for {
		token := lexer.NextToken()
		fmt.Println(token.Type == zeus_lexer.HTMLTagOpen || token.Type == zeus_lexer.HTMLTagClose)
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if token.Type == zeus_lexer.HTMLTagOpen || token.Type == zeus_lexer.HTMLTagClose {
			htmlTags = append(htmlTags, token.Value)
		}
	}

	return htmlTags
}
