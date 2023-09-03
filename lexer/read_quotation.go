package zeus_lexer

func (l *Lexer) readQuotation() Token {
	start := l.Position
	l.Position++
	/*
		for l.Position < len(l.Input) && l.Input[l.Position] != '"' {
			l.Position++
		}

		if l.Position >= len(l.Input) || l.Input[l.Position] != '"' {
			return Token{Type: Unknown, Value: l.Input[start:l.Position]}
		}

		l.Position++
	*/
	return Token{Type: Quotation, Value: l.Input[start:l.Position]}
}
