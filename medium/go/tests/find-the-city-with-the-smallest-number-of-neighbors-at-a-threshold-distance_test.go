package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestFindTheCity(t *testing.T) {
	tests := []struct {
		n         int
		threshold int
		edges     [][]int
		expected  int
	}{
		{
			n:         4,
			edges:     [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}},
			threshold: 4,
			expected:  3, // Example expected output
		},
	}

	for _, test := range tests {
		result := problems.FindTheCity(test.n, test.edges, test.threshold)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
