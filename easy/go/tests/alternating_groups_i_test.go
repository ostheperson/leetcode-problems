package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestNumberOfAlternatingGroups(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{0, 1, 0, 0, 1},
			expected: 3,
		},
		{
			input:    []int{0, 1, 1},
			expected: 1,
		},
	}

	for _, test := range tests {
		result := problems.NumberOfAlternatingGroups(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
