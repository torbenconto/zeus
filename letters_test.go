package zeus_test

import (
	"testing"

	"github.com/torbenconto/zeus"
)

func TestLetterCount(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"Hello, world!", 10},       // Two words with 10 letters in total
		{"The quick brown fox", 16}, // Four words with 16 letters in total
		{"12345", 0},                // No words, should return 0
		{"", 0},                     // Empty input, should return 0
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			actual := zeus.LetterCount(testCase.input)
			if actual != testCase.expected {
				t.Errorf("LetterCount(%q) = %d; expected %d", testCase.input, actual, testCase.expected)
			}
		})
	}
}
