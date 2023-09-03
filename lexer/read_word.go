package zeus_lexer

import "unicode"

func (l *Lexer) readWord() Token {

	start := l.Position
	isAbbreviation := false

	for l.Position < len(l.Input) && (unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-') {
		if l.Input[l.Position] == '.' && l.Position == start+1 {
			isAbbreviation = true
		}
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
