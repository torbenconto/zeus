package zeus_lexer

import "unicode"

func (l *Lexer) ReadApostrophe() Token {
	start := l.Position
	l.Position++

	for l.Position < len(l.Input) && (unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-') {
		l.Position++
	}

	return Token{Type: Word, Value: l.Input[start:l.Position]}
}
