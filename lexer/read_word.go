package zeus_lexer

import (
	"unicode"
)

func (l *Lexer) readWord() Token {
	start := l.Position
	isAbbreviation := false

	// Check if the word starts with a period
	if unicode.IsUpper(rune(l.Input[start:l.Position][0])) && l.Position < len(l.Input) && l.Input[l.Position] == '.' {
		isAbbreviation = true
		l.Position++
	}

	for l.Position < len(l.Input) && (unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-') {
		l.Position++
	}

	if isAbbreviation {
		return Token{
			Type:  Abbreviation,
			Value: l.Input[start:l.Position],
		}
	}

	return Token{
		Type:  Word,
		Value: l.Input[start:l.Position],
	}
}
