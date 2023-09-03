package zeus_lexer

func (l *Lexer) ReadForwardSlash() Token {
	start := l.Position
	l.Position++

	return Token{Type: ForwardSlash, Value: l.Input[start:l.Position]}
}
