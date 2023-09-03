package zeus

import "regexp"

// Extracts a given number of urls from the given string (in) using regex. Returns a list of urls with a max length of (max)
func ExtractUrls(in string, max int) []string {
	regex := regexp.MustCompile(`https?://[^\s]+`)

	urls := regex.FindAllString(in, -1)

	if max > 0 && max < len(urls) {
		urls = urls[:max]
	}

	return urls
}
