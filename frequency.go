package zeus

import "strings"

func FrequencyOfChar(char string, text string) int {
	count := 0
	for _, c := range strings.Split(text, "") {
		if c == char {
			count++
		}
	}
	return count
}
