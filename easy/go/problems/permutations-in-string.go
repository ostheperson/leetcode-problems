package problems

func mapsEqual(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if b[key] != value {
			return false
		}
	}
	return true
}

func CheckInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	patternmap := make(map[rune]int)
	s2map := make(map[rune]int)
	for _, val := range s1 {
		patternmap[val]++
	}
	for end := 0; end < len(s2); end++ {
		if end < len(s1) {
			continue
		}
		for start := 0; start < len(s1); start++ {
			s2map[rune(s2[start+end])]++
		}
		if match := mapsEqual(s2map, patternmap); match {
			return match
		}
	}
	if match := mapsEqual(s2map, patternmap); match {
		return match
	}
	return false
}
