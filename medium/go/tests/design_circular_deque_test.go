package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestdesignCircularDeque(t *testing.T) {
	tests := []struct {
		expected []int
	}{
		{
			expected: []int{},
		},
	}

	for _, test := range tests {
		result := problems.designCircularDeque()
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
