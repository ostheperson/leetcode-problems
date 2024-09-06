package problems

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ContainsDuplicate(nums []int, k int) bool {
	if len(nums) == 1 {
		return false
	}
	hash := make(map[int]int)

	for j := 0; j < len(nums); j++ {
		if prev, ok := hash[nums[j]]; ok {
			if absInt(j-prev) <= k {
				return true
			}
		}
		hash[nums[j]] = j
	}
	return false
}
