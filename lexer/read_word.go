package zeus_lexer

import (
	"regexp"
)

func (l *Lexer) readWord() Token {
	start := l.Position

	// Define a regular expression pattern to match abbreviations
	/*
		http://people.physics.illinois.edu/Celia/Caps&Acronyms.pdf
		"In general, common nouns are not capitalized when they're written out as words, but the abbreviations are ALWAYS capitalized—whether they're units, elements, or acronyms. Elements, even those derived from proper names (curium, francium), are always written lower case when they are written out as words."
		Checking if the word is capitalized and has a period character after it is a not very good, but good enough idea for checking if a word is an abbreviation
	*/
	abbreviationPattern := regexp.MustCompile(`^[A-Za-z]+(\.[A-Za-z]+)*$`)

	for l.Position < len(l.Input) && (isLetterOrSpecialChar(l.Input[l.Position])) {
		l.Position++
	}

	word := l.Input[start:l.Position]

	if abbreviationPattern.MatchString(word) {
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

// Helper function to check if a character is a letter or a special character
func isLetterOrSpecialChar(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '\'' || char == '-'
}
