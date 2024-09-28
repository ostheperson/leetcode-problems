package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems" // Update this path to match your module structure
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			expected: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			input:    "AAAAAAAAAAAAA",
			expected: []string{"AAAAAAAAAA"},
		},
	}

	for _, test := range tests {
		result := problems.FindRepeatedDnaSequences(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
