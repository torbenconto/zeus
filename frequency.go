package zeus

// Returns the number of times a given rune (char) appears in another given string (text)
func FrequencyOfRune(char rune, in string) int {
	frequency := 0
	charRune := char

	for _, c := range in {
		if c == charRune {
			frequency++
		}
	}

	return frequency
}

// Returns the number of times a given string (str) appears in another given string (in)
func FrequencyOfString(str string, in string) int {
	frequency := 0
	words := GetAllWords(in)

	for _, w := range words {
		if w == str {
			frequency++
		}
	}

	return frequency
}
