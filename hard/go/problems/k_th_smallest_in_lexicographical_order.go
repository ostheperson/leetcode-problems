package problems

func FindKthNumber(n int, k int) int {
	var countSteps func(curr int) int
	countSteps = func(curr int) int {
		count := 0
		nei := curr + 1
		for curr <= n {
			count += min(n+1, nei) - curr
			curr *= 10
			nei *= 10
		}
		return count
	}

	curr := 1
	for k != 0 {
		steps := countSteps(curr)
		if k-steps > 0 {
			curr++
			k -= steps
		} else {
			curr *= 10
			k--
		}
	}
	return curr
}
