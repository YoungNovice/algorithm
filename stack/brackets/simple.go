package brackets

import "algorithm/stack"

// 只匹配小括号
// eg: (()(()))
func MatchSmallBrackets(str string) bool {
	stk := &stack.Stack{}
	for _, item := range []rune(str) {
		if item == '(' {
			stk.Push(item)
		} else if item == ')' {
			if stk.Empty() {
				return false
			} else {
				stk.Pop()
			}
		}
	}

	if stk.Empty() {
		return true
	} else {
		return false
	}
}
