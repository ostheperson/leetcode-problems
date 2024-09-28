package problems

import "fmt"

func FindTheLongestSubstring(s string) int {
	// hash := make(map[int]int)
	prefixSum := make([]int, len(s))

	prefix := 0
	for i, ch := range s {
		switch ch {
		case 'a':
			prefix ^= 1 << 0
		case 'e':
			prefix ^= 1 << 1
		case 'i':
			prefix ^= 1 << 2
		case 'o':
			prefix ^= 1 << 3
		case 'u':
			prefix ^= 1 << 4
		}
		prefixSum[i] = prefix
	}
	fmt.Println(prefixSum)
	maxLen := 0

	/*
	  get char,
	  set prefix
	  if in hash
	    set max to lenght of i to hash[prefix]
	  else add to hash

	*/
	// for right, ch := range s {
	// 	if left, exist := hash[prefix]; exist {
	// 		maxLen = right - left + 1
	// 	} else {
	// 		hash[prefix] = right
	// 	}
	// }
	// if prefix == 0 {
	// 	return len(s)
	// }
	return maxLen
}
