package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestFindMinDifference(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input:    []string{"23:59", "00:00"},
			expected: 1,
		},
		{
			input: []string{
				"00:00",
				"03:21",
				"07:30",
				"12:34",
				"16:45",
				"20:15",
				"22:22",
				"23:59",
			},
			expected: 1,
		},
	}

	for _, test := range tests {
		result := problems.FindMinDifference(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
