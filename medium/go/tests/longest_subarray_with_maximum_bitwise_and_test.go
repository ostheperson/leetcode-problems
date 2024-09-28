package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestLongestSubarrayWithMaximumBitwiseAND(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 3, 2, 2},
			expected: 2,
		},
		{
			input: []int{
				311155,
				311155,
				311155,
				311155,
				311155,
				311155,
				311155,
				311155,
				201191,
				311155,
			},
			expected: 8,
		},
	}

	for _, test := range tests {
		result := problems.LongestSubarray(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
