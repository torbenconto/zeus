package zeus

import "strings"

// Runs all sanitization functions on a given string (in). Note that not sanatizing text may result in unexpected results.
func SanitizeAll(in string) string {
	in = ReplaceProprietaryMicrosoftCharacters(in)
	in = ReplaceEmails(in, "", -1)
	in = ReplaceUrls(in, "", -1)
	return in
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

// Replace given amount of emails (max) from given string (in) with a given string (replacer). Calls ExtractEmails(in, max) and then replaces the emails inside the result.
func ReplaceEmails(in string, replacer string, max int) string {
	emails := ExtractEmails(in, max)

	for _, email := range emails {
		in = strings.Replace(in, email, replacer, -1)
	}

	return in
}

// Replace given amount of urls (max) from given string (in) with a given string (replacer). Calls ExtractUrls(in, max) and then replaces the urls inside the result.
func ReplaceUrls(in string, replacer string, max int) string {
	urls := ExtractUrls(in, max)

	for _, url := range urls {
		in = strings.Replace(in, url, replacer, -1)
	}

	return in
}
