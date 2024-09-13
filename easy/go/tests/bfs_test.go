package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems"
)

var tests2d []struct {
	input    [][]int
	expected []int
} = []struct {
	input    [][]int
	expected []int
}{
	{
		input: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		input: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		expected: []int{0, 1, 0},
	},
}

var testsMap []struct {
	input    map[int][]int
	expected []int
} = []struct {
	input    map[int][]int
	expected []int
}{
	{
		input: map[int][]int{
			1: {2, 3},
			2: {1, 4},
			3: {1, 4, 5},
			4: {2, 3},
			5: {3},
		},
		expected: []int{1, 2, 3, 4, 5},
	},
}

func TestBFS2D(t *testing.T) {
	for _, test := range tests2d {
		result := problems.Bfs2D(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}

// func TestBFSMap(t *testing.T) {
// 	for _, test := range testsMap {
// 		result := problems.BfsMap(test.input)
// 		if !reflect.DeepEqual(result, test.expected) {
// 			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
// 		}
// 	}
// }
