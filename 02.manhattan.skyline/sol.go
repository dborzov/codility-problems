package solution

func Solution(H []int) (count int) {
	stack := make([]int, 0, len(H))
	for _, curlev := range H {
		for len(stack) > 0 && stack[len(stack)-1] > curlev {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 && stack[len(stack)-1] == curlev {
			continue
		}
		stack = append(stack, curlev)
		count++
	}
	return
}
