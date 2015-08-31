package solution

type Stack []rune

func (st *Stack) Pop() rune {
	oldStack := *st
	if len(oldStack) == 0 {
		return rune(20)
	}
	le := oldStack[len(oldStack)-1]
	*st = oldStack[:len(oldStack)-1]
	return le
}

func Solution(S string) int {
	stack := make(Stack, 0, len(S))
	for _, sym := range S {
		switch sym {
		case '{':
			stack = append(stack, sym)
		case '[':
			stack = append(stack, sym)
		case '(':
			stack = append(stack, sym)
		case '}':
			if stack.Pop() != '{' {
				return 0
			}
		case ']':
			if stack.Pop() != '[' {
				return 0
			}
		case ')':
			if stack.Pop() != '(' {
				return 0
			}
		}
	}

	if len(stack) != 0 {
		return 0
	}
	return 1
}
