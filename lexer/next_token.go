package zeus_lexer

import (
	"unicode"
)

func (l *Lexer) NextToken() Token {
	// Janky fix for a problem i cannot figure out where it throws an error like this panic: runtime error: index out of range [19] with length 19 [recovered] if there is no punctuation of no whitespace (only sometimes though)
	if len(l.Input) > 0 && l.Input[len(l.Input)-1] != ' ' {
		l.Input += " "
	}
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

	if l.Input[l.Position] == ':' {
		return l.ReadColon()
	}

	return l.ReadPunctuation()
}
