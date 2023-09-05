package zeus

// FrequencyOfRune returns the number of times a given rune (char) appears in another given string (text).
// Parameters:
//   - char: The rune to count in the input string.
//   - in: The input string in which to count the rune.
//
// Returns:
//   - The frequency of the rune in the input string.
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

// PercentFrequencyOfRune calculates the percentage frequency of a specific rune (char) within a string (in).
// Parameters:
//   - char: The rune to calculate the percentage frequency of in the input string.
//   - in: The input string in which to calculate the percentage frequency of the rune.
//
// Returns:
//   - The percentage frequency of the rune in the input string as a float64.
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

// FrequencyOfString returns the number of times a given string (str) appears in another given string (in).
// Parameters:
//   - str: The string to count in the input string.
//   - in: The input string in which to count the string.
//
// Returns:
//   - The frequency of the string in the input string.
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

// PercentFrequencyOfString calculates the percentage frequency of a specific string (str) within another string (in).
// Parameters:
//   - str: The string to calculate the percentage frequency of in the input string.
//   - in: The input string in which to calculate the percentage frequency of the string.
//
// Returns:
//   - The percentage frequency of the string in the input string as a float64.
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
