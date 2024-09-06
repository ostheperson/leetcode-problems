package problems

// import "fmt"

// HashcodePizzaProblem takes parameters and performs its operation.
// Add your implementation here.
func SelectPizzas(maxSlices int, pizzaTypes []int) []int {
	// TODO: Implement
	/*
				  choose largest pizza besed on what slices left
				  from reverse, total_slices += i if slices left is >= i
				  slices_left = 17
				  selected arr []
				  1 iter starting with end
				  if pizzaTypes[3] (8)  <= slices_left  (17) {
				    append pizzaTypes[3]
				    slices_left -= pizzaTypes[3]
				  }
				  2 iter
				  if pizzaTypes[2] (6) <= slices_left (9) {
				    append pizzaTypes[2]
				    slices_left -= pizzaTypes[2]
				  }
		      3 iter
				  if pizzaTypes[1] (5) <= slices_left (3) {
				    append pizzaTypes[1]
				    slices_left -= pizzaTypes[1]
				  }
	*/
	var res []int
	slices_left := maxSlices
	for i := len(pizzaTypes) - 1; i >= 0; i-- {
		if pizzaTypes[i] <= slices_left {
			res = append(res, i)
			slices_left -= pizzaTypes[i]
		}
	}
	reverse(res)
	return res
}

func reverse(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
