package zeus

// FleschKincaidReadingEase calculates the Flesch-Kincaid Reading Ease score for a given text.
// The higher the score, the easier the text is to read.
//
// Parameters:
//   - in: The input text for which to calculate the Reading Ease score.
//
// Returns:
//   - The Flesch-Kincaid Reading Ease score as a float64.
func FleschKincaidReadingEase(in string) float64 {
	return 206.835 - (1.015 * AverageWordsPerSentence(in)) - (84.6 * AverageSyllablesPerWord(in))
}

// FleschKincaidGradeLevel calculates the Flesch-Kincaid Grade Level for a given text.
// It estimates the number of years of education required to understand the text.
//
// Parameters:
//   - in: The input text for which to calculate the Grade Level.
//
// Returns:
//   - The Flesch-Kincaid Grade Level as a float64.
func FleschKincaidGradeLevel(in string) float64 {
	return (0.39 * AverageWordsPerSentence(in)) + (11.8 * AverageSyllablesPerWord(in)) - 15.59
}

// ColemanLiauIndex calculates the Coleman-Liau Index for a given text.
// It estimates the U.S. grade level required to understand the text.
//
// Parameters:
//   - in: The input text for which to calculate the Coleman-Liau Index.
//
// Returns:
//   - The Coleman-Liau Index as a float64.
func ColemanLiauIndex(in string) float64 {
	sentences := SentenceCount(in)
	words := WordCount(in)
	letters := LetterCount(in)

	if sentences == 0 {
		sentences = 1
	}

	return (5.88 * (float64(letters) / float64(words))) - (0.296 * (float64(sentences) / float64(words))) - 15.8
}

// AutomatedReadabilityIndex calculates the Automated Readability Index (ARI) for a given text.
// It estimates the U.S. grade level required to understand the text.
//
// Parameters:
//   - in: The input text for which to calculate the ARI.
//
// Returns:
//   - The Automated Readability Index (ARI) as a float64.
func AutomatedReadabilityIndex(in string) float64 {
	sentences := SentenceCount(in)
	words := WordCount(in)
	letters := LetterCount(in)

	if sentences == 0 {
		sentences = 1
	}

	return (4.71 * (float64(letters) / float64(words))) + (0.5 * (float64(words) / float64(sentences))) - 21.43
}
