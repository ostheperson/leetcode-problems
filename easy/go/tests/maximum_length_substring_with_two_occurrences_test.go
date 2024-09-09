package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestMaximumLengthSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "bbbb",
			expected: 2,
		},
		{
			input:    "bcbbbcba",
			expected: 4,
		},
	}

	for _, test := range tests {
		result := problems.MaximumLengthSubstring(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
