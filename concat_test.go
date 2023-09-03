package zeus_test

import (
	"testing"

	"github.com/torbenconto/zeus"
)

func TestConcat(t *testing.T) {
	tt := []struct {
		input    []string
		expected string
	}{
		{[]string{"Hello, ", "world!"}, "Hello, world!"},
		{[]string{"This", " is", " a", " test"}, "This is a test"},
		{[]string{"One", " ", "two", " ", "three"}, "One two three"},
		{[]string{"", "Empty", ""}, "Empty"},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := zeus.Concat(tc.input...)
			if result != tc.expected {
				t.Errorf("Concat(%v) = %s; expected %s", tc.input, result, tc.expected)
			}
		})
	}
}
