package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestMinimumSubarrayLength(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected int
	}{
		{
			input:    []int{1, 2, 3},
			k:        2,
			expected: 1,
		},
		{
			input:    []int{2, 1, 8},
			k:        10,
			expected: 3,
		},
	}

	for _, test := range tests {
		result := problems.MinimumSubarrayLength(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
