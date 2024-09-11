package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestDecrypt(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{
			input:    []int{5, 7, 1, 4},
			k:        3,
			expected: []int{12, 10, 16, 13},
		},
		{
			input:    []int{2, 4, 9, 3},
			k:        -2,
			expected: []int{12, 5, 6, 13},
		},
	}

	for _, test := range tests {
		result := problems.Decrypt(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
