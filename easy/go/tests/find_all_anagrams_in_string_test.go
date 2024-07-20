package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems" // Update this path to match your module structure
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		expected []int
	}{
		// Add your test cases here
		{
			expected: []int{}, // Example expected output
		},
	}

	for _, test := range tests {
		result := problems.FindAnagrams()
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
