package zeus_lexer

import (
	"unicode"
)

func (l *Lexer) readWord() Token {
	start := l.Position
	isAbbreviation := false

	// Check if the word starts with a period
	if l.Position < len(l.Input) && l.Input[l.Position] == '.' {
		isAbbreviation = true
		l.Position++
	}

	for l.Position < len(l.Input) && (unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-') {
		l.Position++
	}

	word := l.Input[start:l.Position]

	if isAbbreviation && isUppercase(word) {
		return Token{
			Type:  Abbreviation,
			Value: word,
		}
	}

	return Token{
		Type:  Word,
		Value: word,
	}
}

func isUppercase(word string) bool {
	for _, char := range word {
		if !unicode.IsUpper(char) {
			return false
		}
	}
	return true
}
