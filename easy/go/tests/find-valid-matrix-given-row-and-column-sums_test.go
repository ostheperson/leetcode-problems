package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestRestoreMatrix(t *testing.T) {
	tests := []struct {
		rowSum   []int
		colSum   []int
		expected [][]int
	}{
		{
			rowSum:   []int{3, 8},
			colSum:   []int{4, 7},
			expected: [][]int{{3, 0}, {1, 7}},
		},
	}

	for _, test := range tests {
		result := problems.RestoreMatrix(test.rowSum, test.colSum)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("\nrowSum : %v colSum: %v\nexpected :%v\nbut got :%v", test.rowSum, test.colSum, test.expected, result)
		}
	}
}
