package solution

func Solution(A []int) int {
	stack := make([]int, 0, len(A))
	for idx := range A {
		stack = append(stack, idx)
		if len(stack) > 1 && A[stack[len(stack)-1]] != A[stack[len(stack)-2]] {
			stack = stack[:stack[len(stack)-2]]
		}
	}

	if len(stack) > 0 && len(stack) > len(A)/2 {
		return stack[0]
	}

	return -1
}
