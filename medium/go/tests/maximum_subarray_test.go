package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestMaximumSubarray(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			input:    []int{2, -5, 3, 4},
			expected: 7,
		},
	}

	for _, test := range tests {
		result := problems.MaximumSubarray(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
