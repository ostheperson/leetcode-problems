package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestFindTheLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "eleetminicoworoep",
			expected: 13,
		},
	}

	for _, test := range tests {
		result := problems.FindTheLongestSubstring(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
