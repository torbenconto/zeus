package zeus

import "testing"

func TestStringDifference(t *testing.T) {
	tests := []struct {
		a, b, expected string
	}{
		{"AGGTAB", "GXTXAYB", "GTAB"},
		{"abcdef", "xyz", ""},
		{"", "abc", ""},
		{"abc", "abc", "abc"},
		{"a", "b", ""},
		{"aaaa", "aaaa", "aaaa"},
	}

	for _, test := range tests {
		result := StringDifference(test.a, test.b)
		if result != test.expected {
			t.Errorf("StringDifference(%s, %s) = %s; expected %s", test.a, test.b, result, test.expected)
		}
	}
}
