package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestCountKConstraintSubstrings(t *testing.T) {
	tests := []struct {
		input    string
		k        int
		expected int
	}{
		{
			input:    "10101",
			k:        1,
			expected: 12,
		},
		{
			input:    "11111",
			k:        1,
			expected: 15,
		},
	}

	for _, test := range tests {
		result := problems.CountKConstraintSubstrings(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
