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

func PercentFrequencyOfRune(char rune, in string) float64 {
	frequency := 0
	charRune := char

	for _, c := range in {
		if c == charRune {
			frequency++
		}
	}

	return float64(frequency) / float64((len(in) - frequency))
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

func PercentFrequencyOfString(str string, in string) float64 {
	frequency := 0
	words := GetAllWords(in)

	for _, w := range words {
		if w == str {
			frequency++
		}
	}

	return float64(frequency) / float64(len(words)-frequency)
}
