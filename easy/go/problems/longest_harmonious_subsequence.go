package problems

func FindLHS(nums []int) int {
	hash := make(map[int]int)
	for _, val := range nums {
		hash[val]++
	}
	max_len := 0
	for val := range hash {
		if count1, ok := hash[val+1]; ok {
			count0, _ := hash[val]
			lenght := count0 + count1
			max_len = max(max_len, lenght)
		}
	}
	return max_len
}
