package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems" // Update this path to match your module structure
)

func TestSortArrayByIncreasingFrequency(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 1, 2, 2, 2, 3},
			expected: []int{3, 1, 1, 2, 2, 2},
		},
		{
			input:    []int{2, 3, 1, 3, 2},
			expected: []int{1, 3, 3, 2, 2},
		},
	}

	for _, test := range tests {
		result := problems.FrequencySort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
