package problems

func MaximumSubarray(nums []int) int {
	m := nums[0]
	curr := 0
	for _, val := range nums {
		curr = max(curr+val, val)
		m = max(curr, m)
	}
	return m
}
