package zeus

import "regexp"

// Uses regex to extract a given amount of emails (max) from a given string (in). Use a max value of 0 or less to return all emails
func ExtractEmails(in string, max int) []string {
	regex := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`)
	emails := regex.FindAllString(in, -1)

	if max > 0 && max < len(emails) {
		emails = emails[:max]
	}

	return emails
}
