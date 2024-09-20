package problems

func ShortestPalindrome(s string) string {
	r := len(s)
	if r == 0 {
		return ""
	}

	for r > 0 {
		if s[0] == s[r-1] && IsPalindrome(s[:r]) {
			break
		}
		r--
	}
	if r == len(s) {
		return s
	}

	added := []rune(s[r:])
	n := len(added)
	for i := 0; i < n/2; i++ {
		added[i], added[n-1-i] = added[n-1-i], added[i]
	}

	return string(added) + s
}

func IsPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
