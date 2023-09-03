package zeus

import "regexp"

// Uses regex to extract a given amount of emails (max) from a given string (in). Use a max value of 0 or less to return all emails
func ExtractEmails(in string, max int) []string {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	emails := regex.FindAllString(in, -1)

	if max > 0 && max < len(emails) {
		emails = emails[:max]
	}

	return emails
}
