package zeus_lexer

import (
	"fmt"
	"unicode"
)

func (l *Lexer) readWord() Token {
	start := l.Position
	isAbbreviation := false

	for l.Position < len(l.Input) && (unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-' || l.Input[l.Position] == '.') {
		l.Position++
	}

	if l.Position < len(l.Input) && l.Input[l.Position-1] == '.' {
		isAbbreviation = true
	}

	word := l.Input[start:l.Position]
	fmt.Println(word, isAbbreviation, isUppercase(word))

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
