package zeus

import "strings"

// Runs all sanitization functions on a given string (s). Note that not sanatizing text may result in unexpected results.
func SanitizeAll(s string) string {
	return ReplaceProprietaryMicrosoftCharacters(s)
}

// Replaces commonly used (bad) microsoft and google ascii characters with normal characters.
func ReplaceProprietaryMicrosoftCharacters(s string) string {
	replacements := map[string]string{
		"“": "\"",   // Smart quote (left)
		"”": "\"",   // Smart quote (right)
		"™": "(TM)", // Trademark symbol
		"…": "...",  // Ellipsis
		"—": "--",   // Em dash
		"–": "-",    // En dash
		"’": "'",    // weird quote
	}

	output := s
	for microsoftChar, standardChar := range replacements {
		output = strings.ReplaceAll(output, microsoftChar, standardChar)
	}

	return output
}
