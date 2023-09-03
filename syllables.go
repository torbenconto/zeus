package zeus

import (
	"strings"
)

func SyllableCount(word string) (count int) {
	word = strings.ToLower(word)

	count, ok := ProblemWords[word]
	if ok {
		return
	}

	var prefixSuffixCount int
	for _, regex := range PrefixSuffixes[:] {
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

	for _, regex := range SubSyllables[:] {
		if regex.MatchString(word) {
			count--
		}
	}

	for _, regex := range AddSyllables[:] {
		if regex.MatchString(word) {
			count++
		}
	}

	return
}

func TotalSyllableCount(in string) int {
	words := GetAllWords(in)
	syllables := 0

	for _, word := range words {
		syllables += SyllableCount(word)
	}

	return syllables
}

func AverageSyllablesPerWord(in string) float64 {
	return float64(TotalSyllableCount(in)) / float64(WordCount(in))
}
