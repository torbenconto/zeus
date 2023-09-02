package zeus_lexer

func (l *Lexer) ReadPunctuation() Token {

	start := l.Position
	l.Position++

	return Token{
		Type:  Puntuation,
		Value: l.Input[start:l.Position],
	}
}
