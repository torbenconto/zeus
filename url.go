package zeus

import (
	"regexp"
	"strings"
)

// Extracts a given number of urls from the given string (in) using regex. Returns a list of urls with a max length of (max). Use a max value of -1 or less to return every url
func ExtractUrls(in string, max int) []string {
	regex := regexp.MustCompile(`https?://[^\s]+`)

	urls := regex.FindAllString(in, -1)

	if max >= 0 && max < len(urls) {
		urls = urls[:max]
	}

	return urls
}

// Replace given amount of urls (max) from given string (in) with a given string (replacer). Calls ExtractUrls(in, max) and then replaces the urls inside the result.
func ReplaceUrls(in string, replacer string, max int) string {
	urls := ExtractUrls(in, max)

	for _, url := range urls {
		in = strings.Replace(in, url, replacer, -1)
	}

	return in
}
