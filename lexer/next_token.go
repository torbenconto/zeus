package zeus_lexer

import "unicode"

func (l *Lexer) NextToken() Token {
	l.SkipWhiteSpace()

	if l.Position >= len(l.Input) {
		return Token{Type: Unknown, Value: ""}
	}

	if unicode.IsLetter(rune(l.Input[l.Position])) {
		return l.ReadWord()
	}

	if l.Input[l.Position] == '"' {
		return l.ReadQuotation()
	}

	return l.ReadPunctuation()
}
