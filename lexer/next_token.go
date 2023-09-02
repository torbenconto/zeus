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

	if l.Input[l.Position] == '\'' {
		return l.ReadApostrophe()
	}

	if l.Input[l.Position] == ',' {
		return l.ReadComma()
	}

	if l.Input[l.Position] == '(' {
		return l.ReadLeftParenthesis()
	}

	if l.Input[l.Position] == ')' {
		return l.ReadRightParenthesis()
	}

	if unicode.IsDigit(rune(l.Input[l.Position])) {
		return l.ReadNumber()
	}

	return l.ReadPunctuation()
}
