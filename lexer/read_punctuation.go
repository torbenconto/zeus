package zeus_lexer

func (l *Lexer) ReadPunctuation() Token {

	start := l.Position
	for l.Position < len(l.Input) && l.Input[l.Position] == l.Input[start] {
		l.Position++
	}
	return Token{
		Type:  Punctuation,
		Value: l.Input[start:l.Position],
	}
}
