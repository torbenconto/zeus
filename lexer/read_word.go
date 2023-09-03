package zeus_lexer

import (
	"strings"
	"unicode"
)

func (l *Lexer) readWord() Token {
	start := l.Position
	for l.Position < len(l.Input) && (unicode.IsLetter(rune(l.Input[l.Position])) || l.Input[l.Position] == '\'' || l.Input[l.Position] == '-' || l.Input[l.Position] == '.' || unicode.IsDigit(rune(l.Input[l.Position])) || l.Input[l.Position] == '_') {
		l.Position++
	}

	value := l.Input[start:l.Position]
	if strings.HasSuffix(value, ".") && len(value) > 1 && strings.ToUpper(value) != value {
		return Token{Type: Abbreviation, Value: value}
	}

	return Token{
		Type:  Word,
		Value: l.Input[start:l.Position],
	}
}
