package problems

func FindMaxAverage(nums []int, k int) float64 {
	if len(nums) < 1 {
		return 0
	}
	left, current_sum := 0, 0
	max_sum := 0
	for right, val := range nums {
		current_sum += val
		if right < k-1 {
			continue
		}
		max_sum = max(max_sum, current_sum)
		current_sum -= nums[left]
		left++
	}
	return float64(max_sum) / float64(k)
}
