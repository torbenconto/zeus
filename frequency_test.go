package zeus_test

import (
	"testing"

	"github.com/torbenconto/zeus"
)

func TestFrequencyOfChar(t *testing.T) {
	// Test case 1: Single character "a" appears once in "apple"
	result := zeus.FrequencyOfChar("a", "apple")
	if result != 1 {
		t.Errorf(`FrequencyOfChar("a", "apple") = %d; want 1`, result)
	}

	// Test case 2: Single character "z" appears zero times in "apple"
	result = zeus.FrequencyOfChar("z", "apple")
	if result != 0 {
		t.Errorf(`FrequencyOfChar("z", "apple") = %d; want 0`, result)
	}

	// Test case 3: Single character "l" appears three times in "hello"
	result = zeus.FrequencyOfChar("l", "hello")
	if result != 2 {
		t.Errorf(`FrequencyOfChar("l", "hello") = %d; want 2`, result)
	}

	// Test case 4: Single character "A" (uppercase) appears zero times in "apple"
	result = zeus.FrequencyOfChar("A", "apple")
	if result != 0 {
		t.Errorf(`FrequencyOfChar("A", "apple") = %d; want 0`, result)
	}
}

func TestFrequencyOfWord(t *testing.T) {
	tests := []struct {
		word     string
		text     string
		expected int
	}{
		{"banana", "banana apple banana orange banana", 3},
		{"apple", "banana apple banana orange banana", 1},
		{"pear", "banana apple banana orange banana", 0},
		{"", "empty text", 0},
	}

	for _, test := range tests {
		actual := zeus.FrequencyOfWord(test.word, test.text)
		if actual != test.expected {
			t.Errorf("FrequencyOfWord(%s, %s) = %d; expected %d", test.word, test.text, actual, test.expected)
		}
	}
}
