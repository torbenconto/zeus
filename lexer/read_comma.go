package zeus_lexer

func (l *Lexer) ReadComma() Token {
	start := l.Position
	l.Position++
	return Token{Type: Comma, Value: l.Input[start:l.Position]}
}
