package main

import "strings"

func Sanitize(s string) string {
	return ReplaceProprietaryMicrosoftCharacters(s)
}

func ReplaceProprietaryMicrosoftCharacters(s string) string {
	replacements := map[string]string{
		"“": "\"",   // Smart quote (left)
		"”": "\"",   // Smart quote (right)
		"™": "(TM)", // Trademark symbol
		"…": "...",  // Ellipsis
		"—": "--",   // Em dash
		"–": "-",    // En dash
	}

	output := s
	for microsoftChar, standardChar := range replacements {
		output = strings.ReplaceAll(output, microsoftChar, standardChar)
	}

	return output
}
