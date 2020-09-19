package brackets

import "algorithm/stack"

// 匹配三种括号
// eg: asdf(s)(())()){(){}[nihao]}
func MatchComplex(str string) bool {
	stk := &stack.Stack{}

	for _, item := range []rune(str) {
		if item == '(' || item == '[' || item == '{' {
			stk.Push(item)
		} else if item == ')' || item == ']' || item == '}' {
			top, b := stk.Peek()
			if !b {
				return false
			} else if (item == ')' && top == '(') ||
				(item == ']' && top == '[') ||
				(item == '}' && top == '{') {
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
