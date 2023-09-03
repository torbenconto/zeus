package zeus_lexer

func (l *Lexer) readLeftParenthesis() Token {
	start := l.Position
	l.Position++
	return Token{Type: LeftParenthesis, Value: l.Input[start:l.Position]}
}

func (l *Lexer) readRightParenthesis() Token {
	start := l.Position
	l.Position++
	return Token{Type: RightParenthesis, Value: l.Input[start:l.Position]}
}
