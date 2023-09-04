package zeus

import "strings"

// Runs all sanitization functions on a given string (in). Note that not sanatizing text may result in unexpected results.
func SanitizeAll(in string) string {
	in = ReplaceProprietaryMicrosoftCharacters(in)
	in = ReplaceEmails(in, "", -1)
	in = ReplaceUrls(in, "", -1)
	in = ReplaceCommonHTMLCharacters(in)
	return in
}

// Replaces commonly used (bad) microsoft and google ascii characters with normal characters.
func ReplaceProprietaryMicrosoftCharacters(in string) string {
	replacements := map[string]string{
		"“": "\"",   // Smart quote (left)
		"”": "\"",   // Smart quote (right)
		"™": "(TM)", // Trademark symbol
		"…": "...",  // Ellipsis
		"—": "--",   // Em dash
		"–": "-",    // En dash
		"’": "'",    // weird quote
	}

	output := in
	for microsoftChar, standardChar := range replacements {
		output = strings.ReplaceAll(output, microsoftChar, standardChar)
	}

	return output
}

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
