package zeus_lexer

import (
	"fmt"
	"unicode"
)

func (l *Lexer) readWord() Token {
	start := l.Position
	isAbbreviation := false

	for l.Position < len(l.Input) && (unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-' || (l.Input[l.Position] == '.' && !unicode.IsLower(rune(l.Input[l.Position+1])))) {
		fmt.Println(string(l.Input[l.Position]))
		l.Position++
	}

	if l.Position < len(l.Input) && l.Input[l.Position-1] == '.' {
		isAbbreviation = true
	}

	word := l.Input[start:l.Position]

	if !unicode.IsUpper(rune(word[0])) {
		isAbbreviation = false
	}

	if isAbbreviation {
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
