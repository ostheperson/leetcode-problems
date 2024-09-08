package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected float64
	}{
		{
			input:    []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75,
		},
		{
			input:    []int{5},
			k:        1,
			expected: 5.0,
		},
		{
			input:    []int{4, 0, 4, 3, 3},
			k:        5,
			expected: 2.8,
		},
	}

	for _, test := range tests {
		result := problems.FindMaxAverage(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
