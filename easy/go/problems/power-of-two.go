package problems

import "math"

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%2 == 0 {
		return math.Mod(math.Log2(float64(n)), 1) == 0
	}
	return false
}
