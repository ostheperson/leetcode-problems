package problems

func LongestSubarray(nums []int) int {
	k := 0
	for _, val := range nums {
		k = max(val, k)
	}
	maxLen := 0
	curr := 0
	for _, val := range nums {
		if val == k {
			curr += 1
		} else {
			curr = 0
		}
		maxLen = max(maxLen, curr)
	}
	return maxLen
}
