package problems

func MinimumSubarrayLength(nums []int, k int) int {
	b := len(nums) + 1

	for left := range nums {
		c := 0
		for right := left; right < len(nums); right++ {
			c |= nums[right]
			if c >= k {
				b = min(b, right-left+1)
			}
		}
	}
	if b == len(nums)+1 {
		return -1
	}

	return b
}

// improvement. skip inner iteration when there is a found minimum
