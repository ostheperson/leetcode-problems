package problems

import "fmt"

/*
n = 4, edges = [[0,1,3],[1,2,1],[1,3,4],[2,3,1]], distanceThreshold = 4
City 0 -> [City 1, City 2]
City 1 -> [City 0, City 2, City 3]
City 2 -> [City 0, City 1, City 3]
City 3 -> [City 1, City 2]
*/
type Edge struct {
	to     int
	weight int
}

func FindTheCity(n int, edges [][]int, distanceThreshold int) int {
	hash := make(map[int][]Edge)
	for _, val := range edges {
		edge := Edge{
			to:     val[1],
			weight: val[2],
		}
		hash[val[0]] = append(hash[val[0]], edge)
		edge = Edge{
			to:     val[0],
			weight: val[2],
		}

		hash[val[1]] = append(hash[val[1]], edge)
	}
	min_set := false
	minimum := []int{}
	for key := range hash {
		count := getCitiesInThreshold(key, distanceThreshold, hash)
		if min_set == false {
			min_set = true
			minimum = []int{key, count}
		}
		if count < minimum[1] {
			minimum = []int{key, count}
		}
		if count == minimum[1] && key > minimum[0] {
			minimum = []int{key, count}
		}
	}

	return minimum[0]
}

func getCitiesInThreshold(start, threshold int, hash map[int][]Edge) int {
	/*
	  for edge in edges get all edges connected to target
	*/
	visited := make(map[int]bool)
	queue := []int{start}
	count := threshold
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, val := range hash[node] {
			if !visited[val.to] {
				count += val.weight
				visited[val.to] = true
				queue = append(queue, val.to)
			}
		}
	}
	fmt.Println(start)
	fmt.Println(count)
	fmt.Println()

	return 0
}
