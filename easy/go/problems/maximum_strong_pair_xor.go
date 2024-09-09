package problems

import "sort"

func MaximumStrongPairXor(nums []int) int {
	sort.Ints(nums)
	maximum, left, right := 0, 0, 0
	for left != len(nums) {
		if isStrongPair(nums[left], nums[right]) {
			maximum = max(nums[right]^nums[left], maximum)
			right++
			if right == len(nums) {
				left++
				right = left
			}
		} else {
			left++
			right = left
		}
	}
	return maximum
}

func isStrongPair(a, b int) bool {
	return abs(a-b) <= min(a, b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
