package zeus_lexer

import "unicode"

func (l *Lexer) readSymbol() Token {
	start := l.Position
	for l.Position < len(l.Input) && unicode.IsSymbol(rune(l.Input[l.Position])) {
		l.Position++
	}

	return Token{Type: Symbol, Value: l.Input[start:l.Position]}
}
