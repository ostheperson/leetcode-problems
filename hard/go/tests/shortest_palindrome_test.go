package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestShortestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "aacecaaa",
			expected: "aaacecaaa",
		},
	}

	for _, test := range tests {
		result := problems.ShortestPalindrome(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
