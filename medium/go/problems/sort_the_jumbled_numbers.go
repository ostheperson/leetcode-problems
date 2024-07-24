package problems

import (
	"math"
	"sort"
)

func SortTheJumbled(mapping []int, nums []int) []int {
	hash := make(map[int]int)
	for _, x := range nums {
		mappings := 0
		if x == 0 {
			mappings = mapping[x]
		}
		val := x
		count := 0
		for val > 0 {
			last_digit := val % 10
			mappings += mapping[last_digit] * int(math.Pow(float64(10), float64(count)))
			val /= 10
			count++
		}
		hash[x] = mappings

	}
	sort.SliceStable(nums, func(i, j int) bool {
		if hash[nums[i]] == hash[nums[j]] {
			return false
		}
		return hash[nums[i]] < hash[nums[j]]
	})
	return nums
}
