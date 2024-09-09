package problems

func MaximumLengthSubstring(s string) int {
	maximum, count, left, right := 0, 0, 0, 0
	hash := make(map[byte]int)
	if len(s) == 1 {
		return 1
	}

	for right < len(s) {
		val := s[right]
		if freq, exists := hash[val]; exists && freq == 2 {
			count = 0
			left++
			right = left
			hash = make(map[byte]int)
		} else {
			count++
			hash[val]++
			right++
			maximum = max(count, maximum)
		}
	}
	return maximum
}
