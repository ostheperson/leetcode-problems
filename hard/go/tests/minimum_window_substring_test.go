package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected string
	}{
		{
			t:        "BABC",
			s:        "ADOBECODEBANC",
			expected: "BECODEBA",
		},
		// {
		// 	t:        "a",
		// 	s:        "a",
		// 	expected: "a",
		// },
		// {
		// 	t:        "aa",
		// 	s:        "a",
		// 	expected: "",
		// },
	}

	for _, test := range tests {
		result := problems.MinWindow(test.s, test.t)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
