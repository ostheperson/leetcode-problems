package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestFindKthNumber(t *testing.T) {
	tests := []struct {
		input    int
		k        int
		expected int
	}{
		{
			input:    1000000,
			k:        2,
			expected: 10,
		},
	}

	for _, test := range tests {
		result := problems.FindKthNumber(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
