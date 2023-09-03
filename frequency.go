package zeus

func FrequencyOfChar(char string, text string) int {
	count := 0
	charRune := []rune(char)

	for _, c := range text {
		if c == charRune[0] {
			count++
		}
	}

	return count
}
