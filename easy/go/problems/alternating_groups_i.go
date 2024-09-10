package problems

func NumberOfAlternatingGroups(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[(i+1)%len(nums)] {
			if nums[(i+1)%len(nums)] != nums[(i+2)%len(nums)] {
				count++
			}
		}
	}
	return count
}
