package problems

func Bfs2D(nums [][]int) []int {
	x_len, y_len := len(nums), len(nums[0])
	visited := make([][]bool, y_len)
	for i := range visited {
		visited[i] = make([]bool, x_len)
	}
	result := []int{}
	queue := [][2]int{{0, 0}}

	for len(queue) > 0 {
		idx := queue[0]
		queue = queue[1:]

		x := idx[0]
		y := idx[1]

		if !visited[x][y] {
			visited[x][y] = true
			result = append(result, nums[x][y])
		}

		if y > 0 && !visited[x][y-1] {
			queue = append(queue, [2]int{x, y - 1})
		}
		if y < y_len-1 && !visited[x][y+1] {
			queue = append(queue, [2]int{x, y + 1})
		}
		if x < x_len-1 && !visited[x+1][y] {
			queue = append(queue, [2]int{x + 1, y})
		}
		if x > 0 && !visited[x-1][y] {
			queue = append(queue, [2]int{x - 1, y})
		}
	}
	return result
}

func BfsMap(hash map[int][]int) []int {
	visited := make(map[int]bool)
	queue := []int{1}

	for len(queue) > 0 {
		num := queue[0]
		queue = queue[1:]
		if _, exist := visited[num]; !exist {
			visited[num] = true
		}
		for _, i := range hash[num] {
			if _, exist := visited[i]; !exist {
				queue = append(queue, i)
			}
		}
	}
	ans := []int{}
	for i := range visited {
		ans = append(ans, i)
	}
	return ans
}

func Bfs2Drecursion2D(nums [][]int) []int {
	return nil
}

func Bfs2DRecursionMap(hash map[int][]int) []int {
	return nil
}
