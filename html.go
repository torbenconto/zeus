package zeus

import (
	"strings"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

// ExtractHTMLTags extracts a given number of HTML tags from the given string (in) using the lexer.
// It returns a list of HTML tags with a maximum length of (max). Use a max value of -1 or less to return every tag found in the input string.
//
// Parameters:
//   - in: The input string from which to extract HTML tags.
//   - max: The maximum number of HTML tags to extract. Use -1 to extract all HTML tags.
//
// Returns:
//   - A slice of strings containing the extracted HTML tags.
func ExtractHTMLTags(in string, max int) []string {
	lexer := zeus_lexer.NewLexer(in)
	var htmlTags []string

	for {
		token := lexer.NextToken()
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

// ReplaceHTMLTags replaces a given number of HTML tags (max) from the given string (in) with a given string (replacer).
// This function first calls ExtractHTMLTags(in, max) to extract the HTML tags and then replaces them inside the input string with the provided replacer.
//
// Parameters:
//   - in: The input string in which to replace HTML tags.
//   - replacer: The string to replace HTML tags with.
//   - max: The maximum number of HTML tags to replace. Use -1 to replace all HTML tags.
//
// Returns:
//   - The input string with the specified HTML tags replaced by the replacer string.
func ReplaceHTMLTags(in string, replacer string, max int) string {
	tags := ExtractHTMLTags(in, max)

	for _, tag := range tags {
		in = strings.Replace(in, tag, replacer, -1)
	}

	return in
}
