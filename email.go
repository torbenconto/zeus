package zeus

import (
	"regexp"
	"strings"
)

// Uses regex to extract a given amount of emails (max) from a given string (in). Use a max value of -1 or less to return all emails
func ExtractEmails(in string, max int) []string {
	regex := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`)
	emails := regex.FindAllString(in, -1)

	if max >= 0 && max < len(emails) {
		emails = emails[:max]
	}

	return emails
}

// Replace given amount of emails (max) from given string (in) with a given string (replacer). Calls ExtractEmails(in, max) and then replaces the emails inside the result.
func ReplaceEmails(in string, replacer string, max int) string {
	emails := ExtractEmails(in, max)

	for _, email := range emails {
		in = strings.Replace(in, email, replacer, -1)
	}

	return in
}
