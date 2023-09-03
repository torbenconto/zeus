package zeus

func FleschKincaidReadingEase(in string) float64 {
	return 206.835 - (1.015 * AverageWordsPerSentence(in)) - (84.6 * AverageSyllablesPerWord(in))
}

func FleschKincaidGradeLevel(in string) float64 {
	return (0.39 * AverageWordsPerSentence(in)) + (11.8 * AverageSyllablesPerWord(in)) - 15.59
}

func ColemanLiauIndex(in string) float64 {
	sentences := SentenceCount(in)
	words := WordCount(in)
	letters := LetterCount(in)

	if sentences == 0 {
		sentences = 1
	}
	return (5.88 * (float64(letters) / float64(words))) - (0.296 * (float64(sentences) / float64(words))) - 15.8
}

func AutomatedReadabilityIndex(in string) float64 {
	sentences := SentenceCount(in)
	words := WordCount(in)
	letters := LetterCount(in)

	if sentences == 0 {
		sentences = 1
	}

	return (4.71 * (float64(letters) / float64(words))) + (0.5 * (float64(words) / float64(sentences))) - 21.43
}
