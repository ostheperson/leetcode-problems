package tests

import (
	"testing"

	"leetcode-problems/problems" // Update this path to match your module structure
)

func TestHashcodePizzaProblem(t *testing.T) {
	tests := []struct {
		maxSlices  int
		pizzaTypes []int
		expected   []int
	}{
		{17, []int{2, 5, 6, 8}, []int{0, 2, 3}}, // example from the problem
		{15, []int{2, 4, 6, 8}, []int{2, 3}},    // another example
		{10, []int{1, 3, 5, 7}, []int{1, 3}},    // yet another example
	}

	for _, tt := range tests {
		result := problems.SelectPizzas(tt.maxSlices, tt.pizzaTypes)
		if !equal(result, tt.expected, tt.pizzaTypes) {
			t.Errorf("selectPizzas(%d, %v) = %v; expected %v", tt.maxSlices, tt.pizzaTypes, result, tt.expected)
		}
	}
}

func equal(result, expected, pizzaTypes []int) bool {
	sumResult := 0
	for _, idx := range result {
		sumResult += pizzaTypes[idx]
	}

	sumExpected := 0
	for _, idx := range expected {
		sumExpected += pizzaTypes[idx]
	}

	return sumResult == sumExpected
}
