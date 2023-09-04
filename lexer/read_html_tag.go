package zeus_lexer

import "strings"

func (l *Lexer) readHTMLTag() Token {
	start := l.Position

	// Check if it's an opening tag
	if strings.HasPrefix(l.Input[l.Position:], "<") {
		l.Position++ // Move past the '<' character

		// Read until '>' is encountered
		for l.Position < len(l.Input) && l.Input[l.Position] != '>' {
			l.Position++
		}

		if l.Position < len(l.Input) && l.Input[l.Position] == '>' {
			l.Position++
		}

		return Token{Type: HTMLTagOpen, Value: l.Input[start:l.Position]}
	}

	// Check if it's a closing tag
	if strings.HasPrefix(l.Input[l.Position:], "</") {
		l.Position += 2 // Move past the '</' characters

		// Read until '>' is encountered
		for l.Position < len(l.Input) && l.Input[l.Position] != '>' {
			l.Position++
		}

		if l.Position < len(l.Input) && l.Input[l.Position] == '>' {
			l.Position++
		}

		return Token{Type: HTMLTagClose, Value: l.Input[start:l.Position]}
	}

	// If it's not an HTML tag, return an unknown token
	return Token{Type: Unknown, Value: l.Input[start:l.Position]}
}
