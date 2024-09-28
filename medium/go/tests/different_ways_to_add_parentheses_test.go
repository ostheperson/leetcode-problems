package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestDiffWaysToCompute(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{
			input:    "2*3-4*5",
			expected: []int{-34, -14, -10, -10, 10},
		},
	}

	for _, test := range tests {
		result := problems.DiffWaysToCompute(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
