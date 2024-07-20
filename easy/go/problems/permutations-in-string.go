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
	s1map := make(map[rune]int)
	s2map := make(map[rune]int)
	for _, val := range s1 {
		s1map[val]++
	}
	for x := 0; x <= len(s2); x++ {
		if x < len(s1) {
			for y := 0; y < len(s1); y++ {
				s2map[rune(s2[y+x])]++
			}
		} else {
			if match := mapsEqual(s2map, s1map); match {
				return match
			}
		}
	}
	if match := mapsEqual(s2map, s1map); match {
		return match
	}
	return false
}
