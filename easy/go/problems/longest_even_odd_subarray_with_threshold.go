package problems

func LongestAlternatingSubarray(nums []int, threshold int) int {
	max_len, curr_len := 0, 0
	for i, val := range nums {
		if val <= threshold {
			if i == 0 {
				if val%2 == 0 {
					curr_len = 1
				} else {
					curr_len = 0
				}
			} else {
				if nums[i-1] <= threshold && val%2 != nums[i-1]%2 {
					curr_len++
				} else if val%2 == 0 {
					curr_len = 1
				} else {
					curr_len = 0
				}
			}
			max_len = max(max_len, curr_len)
		} else {
			curr_len = 0
		}
	}
	return max_len
}
