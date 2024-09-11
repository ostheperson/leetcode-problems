package problems

func CountKConstraintSubstrings(s string, k int) int {
	count := 0
	for left := range s {
		count0 := 0
		count1 := 0
		for right := left; right < len(s); right++ {
			if s[right] == '1' {
				count1++
			} else {
				count0++
			}
			if count0 > k && count1 > k {
				break
			} else {
				count++
			}
		}
	}
	return count
}
