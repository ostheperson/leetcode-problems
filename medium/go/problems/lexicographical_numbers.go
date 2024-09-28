package problems

func LexicalOrder(n int) []int {
	ans := []int{}
	var dfs func(curr int)
	dfs = func(curr int) {
		if curr > n {
			return
		}
		ans = append(ans, curr)
		for i := 0; i <= 9; i++ {
			next := 10*curr + i
			if next > n {
				break
			}
			dfs(next)
		}
	}

	for i := 1; i <= 9; i++ {
		if i > n {
			break
		}
		dfs(i)
	}
	return ans
}
