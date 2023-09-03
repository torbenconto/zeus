package zeus_lexer

import "unicode"

func (l *Lexer) readNumber() Token {
	start := l.Position
	for l.Position < len(l.Input) && unicode.IsDigit(rune(l.Input[l.Position])) {
		l.Position++
	}

	return Token{Type: Number, Value: l.Input[start:l.Position]}
}
