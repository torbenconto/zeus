package zeus_lexer

import (
	"unicode"
)

func (l *Lexer) readWord() Token {
	start := l.Position

	for l.Position < len(l.Input) && (unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-') {
		l.Position++
	}

	if isAbbreviation(l.Input[start:l.Position]) {
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

func isAbbreviation(word string) bool {
	for _, char := range word {
		if !unicode.IsUpper(char) {
			return false
		}
	}
	return true
}
