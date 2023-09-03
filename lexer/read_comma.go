package zeus_lexer

func (l *Lexer) readComma() Token {
	start := l.Position
	l.Position++
	return Token{Type: Comma, Value: l.Input[start:l.Position]}
}
