package main

import (
	"fmt"

	zeus_lexer "github.com/torbenconto/zeus/lexer"
)

func main() {
	input := `Since those days, a number of revolutions in programming have occurred. One
	revolution with which we are all very familiar is the revolution of languages. First, in
	the late 1940s, came assemblers. These "languages” relieved the programmers of the
	drudgery of translating their programs into binary. In 1951, Grace Hopper invented
	A0, the first compiler. In fact, she coined the term compiler. Fortran was invented in
	1953 (the year after I was born). What followed was an unceasing flood of new
	programming languages—COBOL, PL/1, SNOBOL, C, Pascal, C++, Java, ad
	infinitum.
	Another, probably more significant, revolution was in programming paradigms.
	Paradigms are ways of programming, relatively unrelated to languages. A paradigm
	tells you which programming structures to use, and when to use them. To date, there
	have been three such paradigms.`

	lexer := zeus_lexer.NewLexer(Sanitize(input))
	sentenceCount := 0

	for {
		token := lexer.NextToken()
		if token.Type == zeus_lexer.Unknown {
			break
		}

		if token.Type == zeus_lexer.Punctuation {
			sentenceCount++
		}

		fmt.Printf("Type: %v, Value: %s\n", token.Type, token.Value)
	}

	fmt.Printf("Total Sentences: %d\n", sentenceCount)
}
