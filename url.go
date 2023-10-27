package zeus

import (
	"regexp"
	"strings"
)

const UrlRegex = `https?://[^\s]+`

// ExtractUrls extracts a given number of URLs from the given string (in) using regular expressions.
// It returns a list of URLs with a maximum length of (max). Use a max value of -1 or less to return every URL found in the input string.
//
// Parameters:
//   - in: The input string from which to extract URLs.
//   - max: The maximum number of URLs to extract. Use -1 to extract all URLs.
//
// Returns:
//   - A slice of strings containing the extracted URLs.
func ExtractUrls(in string, max int) []string {
	regex := regexp.MustCompile(UrlRegex)

	urls := regex.FindAllString(in, -1)

	if max >= 0 && max < len(urls) {
		urls = urls[:max]
	}

	return urls
}

// Returns a bool checking if the given text matches a url regex
func IsUrl(in string) bool {
	regex := regexp.MustCompile(UrlRegex)

	return regex.MatchString(in)
}

// ReplaceUrls replaces a given number of URLs (max) from the given string (in) with a given string (replacer).
// This function first calls ExtractUrls(in, max) to extract the URLs and then replaces them inside the input string with the provided replacer.
//
// Parameters:
//   - in: The input string in which to replace URLs.
//   - replacer: The string to replace URLs with.
//   - max: The maximum number of URLs to replace. Use -1 to replace all URLs.
//
// Returns:
//   - The input string with the specified URLs replaced by the replacer string.
func ReplaceUrls(in string, replacer string, max int) string {
	urls := ExtractUrls(in, max)

	for _, url := range urls {
		in = strings.Replace(in, url, replacer, -1)
	}

	return in
}
