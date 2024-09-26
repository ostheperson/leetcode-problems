package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestSumPrefixScores(t *testing.T) {
	tests := []struct {
		input    []string
		expected []int
	}{
		{
			input:    []string{"abc", "ab", "bc", "b"},
			expected: []int{5, 4, 3, 2},
		},
	}

	for _, test := range tests {
		result := problems.SumPrefixScores(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
