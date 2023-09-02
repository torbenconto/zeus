package zeus_lexer

import "unicode"

func (l *Lexer) SkipWhiteSpace() {
	for l.Position < len(l.Input) && unicode.IsSpace(rune(l.Input[l.Position])) {
		l.Position++
	}
}
