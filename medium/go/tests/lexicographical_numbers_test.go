package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestLexicographicalNumbers(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{
			input:    13,
			expected: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, test := range tests {
		result := problems.LexicalOrder(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
