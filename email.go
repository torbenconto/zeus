package zeus

import (
	"regexp"
	"strings"
)

// ExtractEmails uses regular expressions to extract a given amount of emails (max) from a given string (in).
// Use a max value of -1 or less to return all emails found in the input string.
//
// Parameters:
//   - in: The input string from which to extract emails.
//   - max: The maximum number of emails to extract. Use -1 to extract all emails.
//
// Returns:
//   - A slice of strings containing the extracted emails.
func ExtractEmails(in string, max int) []string {
	regex := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`)
	emails := regex.FindAllString(in, -1)

	if max >= 0 && max < len(emails) {
		emails = emails[:max]
	}

	return emails
}

// ReplaceEmails replaces a given amount of emails (max) from a given string (in) with a given string (replacer).
// This function first calls ExtractEmails(in, max) to extract the emails and then replaces them inside the input string with the provided replacer.
//
// Parameters:
//   - in: The input string in which to replace emails.
//   - replacer: The string to replace emails with.
//   - max: The maximum number of emails to replace. Use -1 to replace all emails.
//
// Returns:
//   - The input string with the specified emails replaced by the replacer string.
func ReplaceEmails(in string, replacer string, max int) string {
	emails := ExtractEmails(in, max)

	for _, email := range emails {
		in = strings.Replace(in, email, replacer, -1)
	}

	return in
}
