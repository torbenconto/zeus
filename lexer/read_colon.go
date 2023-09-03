package zeus_lexer

func (l *Lexer) ReadColon() Token {
	start := l.Position
	l.Position++
	return Token{Type: Colon, Value: l.Input[start:l.Position]}
}
