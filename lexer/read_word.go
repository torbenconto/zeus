package zeus_lexer

import (
	"fmt"
	"unicode"
)

func (l *Lexer) readWord() Token {
	start := l.Position

	for l.Position < len(l.Input) && (unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-' || l.Input[l.Position] == '.') {
		l.Position++
	}

	/*
		http://people.physics.illinois.edu/Celia/Caps&Acronyms.pdf
		"In general, common nouns are not capitalized when they're written out as words, but the abbreviations are ALWAYS capitalized—whether they're units, elements, or acronyms. Elements, even those derived from proper names (curium, francium), are always written lower case when they are written out as words."
		Checking if the word is capitalized and has a period character after it is a not very good, but good enough idea for checking if a word is an a
	*/
	if isAbbreviation(l.Input[start:l.Position]) {
		fmt.Println(string(l.Input[start:l.Position]))
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
