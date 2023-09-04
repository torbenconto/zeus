package zeus

import (
	"fmt"
	"strings"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

// Extracts a given number of html tags from the given string (in) using the lexer. Returns a list of html tags with a max length of (max). Use a max value of -1 or less to return every tag
func ExtractHTMLTags(in string, max int) []string {
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

	if max >= 0 && max < len(htmlTags) {
		htmlTags = htmlTags[:max]
	}

	return htmlTags
}

// Replace given amount of Html Tags (max) from given string (in) with a given string (replacer). Calls ExtractHTMLTags(in, max) and then replaces the tags inside the result.
func ReplaceHTMLTags(in string, replacer string, max int) string {
	tags := ExtractHTMLTags(in, max)

	for _, tag := range tags {
		in = strings.Replace(in, tag, replacer, -1)
	}

	return in
}
