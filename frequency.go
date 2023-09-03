package zeus

// Returns the number of times a given rune (char) appears in another given string (text)
func FrequencyOfRune(char rune, text string) int {
	frequency := 0
	charRune := char

	for _, c := range text {
		if c == charRune {
			frequency++
		}
	}

	return frequency
}

// Returns the number of times a given string (word) appears in another given string (text)
func FrequencyOfWord(word string, text string) int {
	frequency := 0
	words := GetAllWords(text)

	for _, w := range words {
		if w == word {
			frequency++
		}
	}

	return frequency
}
