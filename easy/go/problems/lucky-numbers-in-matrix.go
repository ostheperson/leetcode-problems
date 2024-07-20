package problems

func LuckyNumbers(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	minMap := make(map[int]bool)

	for i := 0; i < len(matrix); i++ {
		var minx int = matrix[i][0]
		for j := 0; j < len(matrix[0]); j++ {
			minx = min(matrix[i][j], minx)
		}
		minMap[minx] = true
	}
	for i := 0; i < len(matrix[0]); i++ {
		var maxy int = 0
		for j := 0; j < len(matrix); j++ {
			maxy = max(matrix[j][i], maxy)
		}
		if _, exists := minMap[maxy]; exists {
			return []int{maxy}
		}
	}

	return []int{}
}
