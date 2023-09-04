package zeus_lexer

import "unicode"

func (l *Lexer) readHTMLTag() Token {
	start := l.Position

	l.Position++

	if l.Position < len(l.Input) && l.Input[l.Position] == '/' {
		l.Position++
		for l.Position < len(l.Input) && l.Input[l.Position] != '>' && !unicode.IsSpace(rune(l.Input[l.Position])) {
			l.Position++
		}
		if l.Position < len(l.Input) && l.Input[l.Position] == '>' {
			l.Position++
		} else {
			return Token{Type: Unknown, Value: l.Input[start:l.Position]}
		}
		return Token{Type: HTMLTagClose, Value: l.Input[start:l.Position]}
	}

	for l.Position < len(l.Input) && l.Input[l.Position] != '>' && !unicode.IsSpace(rune(l.Input[l.Position])) {
		l.Position++
	}

	if l.Position < len(l.Input) && l.Input[l.Position] == '>' {
		l.Position++
	} else {
		return Token{Type: Unknown, Value: l.Input[start:l.Position]}
	}

	return Token{Type: HTMLTagOpen, Value: l.Input[start:l.Position]}
}
