package main

import "fmt"

// NextGreater returns, for each element, the first strictly greater element
// to its right, or -1. A monotonic stack of indexes still waiting for their
// answer makes the whole pass linear.
func NextGreater(nums []int) []int {
	out := make([]int, len(nums))
	stack := []int{} // indexes whose next-greater is not yet known
	for i, v := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
			out[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for _, i := range stack {
		out[i] = -1
	}
	return out
}

func main() {
	fmt.Println(NextGreater([]int{6, 8, 0, 1, 3}))
}
