package zeus

import "strings"

// SanitizeAll runs all sanitization functions on a given string (in).
// It performs multiple text sanitization operations to clean and normalize the input string.
//
// Parameters:
//   - in: The input string to sanitize.
//
// Returns:
//   - The sanitized string.
func SanitizeAll(in string) string {
	in = ReplaceProprietaryMicrosoftCharacters(in)
	in = ReplaceEmails(in, "", -1)
	in = ReplaceUrls(in, "", -1)
	in = ReplaceCommonHTMLCharacters(in)
	in = RemoveExtraWhitespace(in)
	return in
}

// ReplaceProprietaryMicrosoftCharacters replaces commonly used (bad) Microsoft and Google ASCII characters with normal characters.
//
// Parameters:
//   - in: The input string to perform replacements on.
//
// Returns:
//   - The input string with proprietary Microsoft characters replaced.
func ReplaceProprietaryMicrosoftCharacters(in string) string {
	replacements := map[string]string{
		"“": "\"",   // Smart quote (left)
		"”": "\"",   // Smart quote (right)
		"™": "(TM)", // Trademark symbol
		"…": "...",  // Ellipsis
		"—": "--",   // Em dash
		"–": "-",    // En dash
		"’": "'",    // Weird quote
	}

	output := in
	for microsoftChar, standardChar := range replacements {
		output = strings.ReplaceAll(output, microsoftChar, standardChar)
	}

	return output
}

// ReplaceCommonHTMLCharacters replaces commonly used HTML entities with their corresponding characters.
//
// Parameters:
//   - in: The input string to perform replacements on.
//
// Returns:
//   - The input string with common HTML characters replaced.
func ReplaceCommonHTMLCharacters(in string) string {
	replacements := map[string]string{
		"&lt;":   "<",   // Less than sign
		"&gt;":   ">",   // Greater than sign
		"&amp;":  "&",   // Ampersand
		"&copy;": "(c)", // Copyright symbol
		"&reg;":  "(R)", // Registered trademark symbol
		"&nbsp;": " ",   // Non-breaking space
	}

	output := in
	for char, replacement := range replacements {
		output = strings.ReplaceAll(output, char, replacement)
	}

	return output
}
