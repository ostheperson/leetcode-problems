package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

func TestLuckyNumbers(t *testing.T) {
	tests := []struct {
		expected []int
	}{}

	for _, test := range tests {
		result := problems.LuckyNumbers()
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
