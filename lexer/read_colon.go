package zeus_lexer

func (l *Lexer) readColon() Token {
	start := l.Position
	l.Position++
	return Token{Type: Colon, Value: l.Input[start:l.Position]}
}
