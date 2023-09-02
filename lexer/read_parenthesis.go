package zeus_lexer

func (l *Lexer) ReadLeftParenthesis() Token {
	start := l.Position
	l.Position++
	return Token{Type: LeftParenthesis, Value: l.Input[start:l.Position]}
}

func (l *Lexer) ReadRightParenthesis() Token {
	start := l.Position
	l.Position++
	return Token{Type: RightParenthesis, Value: l.Input[start:l.Position]}
}
