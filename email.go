package zeus

import "regexp"

// Uses regex to extract a given amount of emails (max) from a given string (in). Use a max value of 0 or less to return all emails
func ExtractEmails(in string, max int) []string {
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	emails := regex.FindAllString(in, -1)

	if max > 0 && max < len(emails) {
		emails = emails[:max]
	}

	return emails
}
