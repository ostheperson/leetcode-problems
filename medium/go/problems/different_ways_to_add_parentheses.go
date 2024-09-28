package problems

import "strconv"

func DiffWaysToCompute(expression string) []int {
	result := []int{}
	for i, ch := range expression {
		if ch == '+' || ch == '*' || ch == '-' {
			left := DiffWaysToCompute(expression[:i])
			right := DiffWaysToCompute(expression[i+1:])

			for _, a := range left {
				for _, b := range right {
					switch ch {
					case '+':
						result = append(result, a+b)
					case '-':
						result = append(result, a-b)
					case '*':
						result = append(result, a*b)
					}
				}
			}
		}
	}
	if len(result) == 0 {
		num, _ := strconv.Atoi(expression)
		result = append(result, num)
	}
	return result
}
