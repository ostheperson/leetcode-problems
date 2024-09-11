package problems

func Decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	for i := range code {
		res[i] = calcValue(i, k, code)
	}
	return res
}

func calcValue(i, k int, nums []int) int {
	res := 0
	if k > 0 {
		j := 1
		for j <= k {
			res += nums[(i+j)%len(nums)]
			j++
		}
	}
	if k < 0 {
		j := 1
		for j <= abs(k) {
			res += nums[(i+len(nums)-j)%len(nums)]
			j++
		}
	}
	return res
}
