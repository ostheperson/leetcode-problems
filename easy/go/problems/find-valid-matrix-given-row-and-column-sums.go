package problems

func RestoreMatrix(rowSum []int, colSum []int) [][]int {
	if len(rowSum) == 1 && len(colSum) == 1 {
		return [][]int{{rowSum[0]}}
	}
	a := len(rowSum)
	b := len(colSum)
	matrix := make([][]int, a)
	for i := range matrix {
		matrix[i] = make([]int, b)
	}

	for rowIdx, row := range matrix {
		for colIdx := range row {
			value := min(rowSum[rowIdx], colSum[colIdx])
			matrix[rowIdx][colIdx] = value
			rowSum[rowIdx] -= value
			colSum[colIdx] -= value
		}
	}
	return matrix
}
