package zeus_lexer

import (
	"fmt"
	"unicode"
)

func (l *Lexer) readWord() Token {
	start := l.Position
	isAbbreviation := false

	for l.Position < len(l.Input) && (unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-' || l.Input[l.Position] == '.') {
		if l.Input[l.Position] == '.' {
			isAbbreviation = true
		}
		l.Position++
	}

	word := l.Input[start:l.Position]

	// Check if the last character is a period and remove it from the word
	if isAbbreviation && len(word) > 0 && word[len(word)-1] == '.' {
		word = word[:len(word)-1]
	}

	if isAbbreviation {
		fmt.Println(word, isAbbreviation, unicode.IsUpper(rune(word[0])))
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
