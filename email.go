package zeus

import "regexp"

func ExtractEmails(in string, max int) []string {
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	emails := regex.FindAllString(in, -1)

	if max > 0 && max < len(emails) {
		emails = emails[:max]
	}

	return emails
}
