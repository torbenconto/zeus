package zeus_lexer

import "unicode"

// ReadQuotation reads and returns a quotation token.
func (l *Lexer) ReadQuotation() Token {
	start := l.Position
	l.Position++

	for l.Position < len(l.Input) && l.Input[l.Position] != '"' {
		// Check for words inside the quotation.
		if unicode.IsLetter(rune(l.Input[l.Position])) {
			// Create a new lexer context for quoted text.
			quotedLexer := &Lexer{Input: l.Input, Position: l.Position}
			quotedLexer.SkipWhiteSpace()

			// Read and ignore words inside the quotation.
			for {
				quotedToken := quotedLexer.NextToken()
				if quotedToken.Type == Unknown {
					break
				}
				if quotedToken.Type == Word {
					// Ignore words inside the quotation.
				}
			}

			// Update the main lexer position to the end of the quoted text.
			l.Position = quotedLexer.Position
		}

		l.Position++
	}

	if l.Position >= len(l.Input) || l.Input[l.Position] != '"' {
		return Token{Type: Unknown, Value: l.Input[start:l.Position]}
	}

	l.Position++

	return Token{Type: Quotation, Value: l.Input[start:l.Position]}
}
