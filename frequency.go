package zeus

func FrequencyOfChar(char string, text string) int {
	frequency := 0
	charRune := []rune(char)

	for _, c := range text {
		if c == charRune[0] {
			frequency++
		}
	}

	return frequency
}

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
