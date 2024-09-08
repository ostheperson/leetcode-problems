package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestmaximumStrongPairXor(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: 7,
		},
		{
			input:    []int{10, 100},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := problems.MaximumStrongPairXor(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
