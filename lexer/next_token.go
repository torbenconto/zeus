package zeus_lexer

import (
	"unicode"
)

func (l *Lexer) NextToken() Token {
	// Janky fix for a problem i cannot figure out where it throws an error like this panic: runtime error: index out of range [19] with length 19 [recovered] if there is no punctuation of no whitespace (only sometimes though)
	if len(l.Input) > 0 && l.Input[len(l.Input)-1] != ' ' {
		l.Input += " "
	}
	l.skipWhiteSpace()

	if l.Position >= len(l.Input) {
		return Token{Type: Unknown, Value: ""}
	}

	if unicode.IsLetter(rune(l.Input[l.Position])) {
		return l.readWord()
	}

	if l.Input[l.Position] == '"' {
		return l.readQuotation()
	}

	//                                this is bad because "-" can also be used in lists but for now we ignore that bcs i am lazy
	if l.Input[l.Position] == '\'' || l.Input[l.Position] == '-' {
		return l.readApostrophe()
	}

	if l.Input[l.Position] == ',' {
		return l.readComma()
	}

	if l.Input[l.Position] == '(' {
		return l.readLeftParenthesis()
	}

	if l.Input[l.Position] == ')' {
		return l.readRightParenthesis()
	}

	if unicode.IsDigit(rune(l.Input[l.Position])) {
		return l.readNumber()
	}

	if unicode.IsSymbol(rune(l.Input[l.Position])) {
		return l.readSymbol()
	}

	if l.Input[l.Position] == '/' {
		return l.readForwardSlash()
	}

	if l.Input[l.Position] == ':' {
		return l.readColon()
	}

	if l.Input[l.Position] == '<' {
		return l.readHTMLTag()
	}

	// TODO: handle urls and abbreviations

	return l.readPunctuation()
}
