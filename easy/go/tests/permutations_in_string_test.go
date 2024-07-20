package tests

import (
	"testing"

	"leetcode-problems/problems"
)

func TestArePermutations(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"abc", "bca", true},
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		// {"", "", true},
		// {"aabbcc", "ccbbaa", true},
		// {"abc", "ab", false},
		// {"123", "321", true},
		// {"", "a", false},
	}

	for _, test := range tests {
		result := problems.CheckInclusion(test.str1, test.str2)
		if result != test.expected {
			t.Errorf("CheckInclusions(%q, %q) = %v; expected %v", test.str1, test.str2, result, test.expected)
		}
	}
}
