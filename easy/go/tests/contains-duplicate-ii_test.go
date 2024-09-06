package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems" // Update this path to match your module structure
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			input:    []int{4, 1, 2, 3, 1, 5},
			k:        3,
			expected: true,
		},
	}

	for _, test := range tests {
		result := problems.ContainsDuplicate(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
