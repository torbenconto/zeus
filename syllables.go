package zeus

import (
	"strings"
)

// Returns the amount of syllables in the given word.
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

// Iterates over every word in the given string (in) and returns the amount of syllables in each word added up.
func TotalSyllableCount(in string) int {
	words := GetAllWords(in)
	syllables := 0

	for _, word := range words {
		syllables += SyllableCount(word)
	}

	return syllables
}

// Returns the amount of syllables in the given string (in) over the amount of words in the given string (in)
func AverageSyllablesPerWord(in string) float64 {
	return float64(TotalSyllableCount(in)) / float64(WordCount(in))
}
