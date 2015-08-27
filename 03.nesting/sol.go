package solution

func Solution(S string) int {
	stackIdx := 0
	for _, sym := range S {
		if sym == '(' {
			stackIdx++
		}
		if sym == ')' {
			stackIdx--
			if stackIdx < 0 {
				return 0
			}
		}
	}
	if stackIdx != 0 {
		return 0
	}
	return 1
}
