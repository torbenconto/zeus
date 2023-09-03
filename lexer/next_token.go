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

	//                                this is bad because "-" can also be used in lists but for now we ignore that bcs i am lazy
	if l.Input[l.Position] == '\'' || l.Input[l.Position] == '-' {
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

	if unicode.IsSymbol(rune(l.Input[l.Position])) {
		return l.ReadSymbol()
	}

	if l.Input[l.Position] == '/' {
		return l.ReadForwardSlash()
	}

	return l.ReadPunctuation()
}
