package zeus

import (
	"strings"
)

// SyllableCount returns the number of syllables in the given word.
//
// Parameters:
//   - word: The word for which to count syllables.
//
// Returns:
//   - The number of syllables in the word.
func SyllableCount(word string) (count int) {
	word = strings.ToLower(word)

	count, ok := problemWords[word]
	if ok {
		return
	}

	var prefixSuffixCount int
	for _, regex := range prefixSuffixes[:] {
		if regex.MatchString(word) {
			word = regex.ReplaceAllString(word, "")
			prefixSuffixCount++
		}
	}

	var wordPartCount int
	for _, wordPart := range consonantsRegexp.Split(word, -1) {
		if len(wordPart) > 0 {
			wordPartCount++
		}
	}

	count = wordPartCount + prefixSuffixCount

	for _, regex := range subSyllables[:] {
		if regex.MatchString(word) {
			count--
		}
	}

	for _, regex := range addSyllables[:] {
		if regex.MatchString(word) {
			count++
		}
	}

	return
}

// TotalSyllableCount iterates over every word in the given string (in) and returns the total number of syllables in all words combined.
//
// Parameters:
//   - in: The input string for which to calculate the total syllable count.
//
// Returns:
//   - The total number of syllables in the input string.
func TotalSyllableCount(in string) int {
	words := GetAllWords(in)
	syllables := 0

	for _, word := range words {
		syllables += SyllableCount(word)
	}

	return syllables
}

// AverageSyllablesPerWord calculates the average number of syllables per word in the given string (in).
//
// Parameters:
//   - in: The input string for which to calculate the average syllables per word.
//
// Returns:
//   - The average number of syllables per word as a float64.
func AverageSyllablesPerWord(in string) float64 {
	return float64(TotalSyllableCount(in)) / float64(WordCount(in))
}
