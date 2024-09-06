package tests

import (
	"fmt"
	"reflect"
	"testing"

	"leetcode-problems/problems" // Update this path to match your module structure
)

func TestFindLHS(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 3, 2, 2, 5, 2, 3, 7},
			expected: 5,
		},
		{
			input:    []int{1, 2, 2, 1},
			expected: 4,
		},
		{
			input:    []int{1, 3, 5, 7, 9, 11, 13, 15, 17},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := problems.FindLHS(test.input)
		fmt.Println(result)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
