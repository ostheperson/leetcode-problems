package problems

import (
	"fmt"
	"sort"
)

func FrequencySort(nums []int) []int {
	hash := make(map[int]int)
	orderedArray := []int{}
	var result []int

	for _, val := range nums {
		if _, exists := hash[val]; exists {
			hash[val] += 1
		} else {
			hash[val] += 1
			orderedArray = append(orderedArray, val)
		}
	}

	sort.Slice(orderedArray, func(i, j int) bool {
		return hash[orderedArray[i]] < hash[orderedArray[j]]
	})

	fmt.Println(orderedArray)
	for _, element := range orderedArray {
		count := hash[element]
		for i := 0; i < count; i++ {
			result = append(result, element)
		}
	}
	return nums
}
