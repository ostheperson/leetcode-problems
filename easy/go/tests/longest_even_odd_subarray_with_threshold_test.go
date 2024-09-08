package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestLongestAlternatingSubarray(t *testing.T) {
	tests := []struct {
		input     []int
		threshold int
		expected  int
	}{
		{
			input:     []int{3, 2, 5, 4},
			threshold: 5,
			expected:  3,
		},
		{
			input:     []int{1, 10, 5},
			threshold: 7,
			expected:  0,
		},
	}

	for _, test := range tests {
		result := problems.LongestAlternatingSubarray(test.input, test.threshold)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
