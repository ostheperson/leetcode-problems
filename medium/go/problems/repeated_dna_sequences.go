package problems

func FindRepeatedDnaSequences(s string) []string {
	hash := make(map[string]int)
	if len(s) < 10 {
		return []string{}
	}
	times := len(s) - 10
	for i := 0; i <= times; i++ {
		hash[s[i:i+10]]++
	}
	var res []string
	for key, val := range hash {
		if val > 1 {
			res = append(res, key)
		}
	}
	return res
}
