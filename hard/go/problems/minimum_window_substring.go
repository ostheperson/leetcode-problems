package problems

func MinWindow(s, t string) string {
	tMap := make(map[rune]int)
	for _, i := range t {
		tMap[i] += 1
	}
	left := 0
	sMap := make(map[rune]int)

	minLeft, minRight := 0, len(s)
	have, need := 0, len(tMap)
	resLen := len(s) + 1

	for right := range s {
		val := rune(s[right])
		if _, exist := tMap[val]; exist {
			sMap[val]++
			if tMap[val] == sMap[val] {
				have++
			}
		}
		for have == need {
			if right-left+1 < resLen {
				minLeft, minRight = left, right
				resLen = minRight - minLeft + 1
			}

			last := rune(s[left])
			if _, exist := tMap[last]; exist {
				sMap[last]--
				if tMap[last] > sMap[last] {
					have--
				}
			}
			left++
		}
	}
	if resLen == len(s)+1 {
		return ""
	}
	return s[minLeft : minRight+1]
}
