package zeus_lexer

import "unicode"

func (l *Lexer) ReadWord() Token {

	start := l.Position
	for l.Position < len(l.Input) && unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-' {
		l.Position++
	}

	return Token{
		Type:  Word,
		Value: l.Input[start:l.Position],
	}
}
