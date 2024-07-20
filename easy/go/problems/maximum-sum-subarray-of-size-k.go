package problems

import "fmt"

func maxSumOfContiguousSubArray(arr []int, k int) int {
	var maxSum, start, currentSum int = 0, 0, 0
	for end := 0; end < len(arr); end++ {
		currentSum += arr[end]
		if end >= k-1 {
			if currentSum > maxSum {
				maxSum = currentSum
			}
			currentSum -= arr[start]
			start++
		}
	}

	return maxSum
}

func main() {
	arr := []int{2, 1, 5, 1, 3, 20}
	k := 3
	fmt.Println(maxSumOfContiguousSubArray(arr, k))
}
