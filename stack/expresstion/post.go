package expresstion

import "algorithm/stack"

// 后缀表达式求值
func PostfixExpression(str string) int {
	s := &stack.Stack{}
	temp := 0
	runes := []rune(str)
	for i, item := range runes {
		if item >= '0' && item <= '9' {
			temp = temp * 10 + int(item) - '0'
			if runes[i+1] == ' ' {
				s.Push(temp)
				temp = 0
			}
			continue
		}
		if item == '+' {
			x, _ := s.Pop()
			y, _ := s.Pop()
			xInt, _ := x.(int)
			yInt, _ := y.(int)
			s.Push(yInt + xInt)
		} else if item == '-' {
			x, _ := s.Pop()
			y, _ := s.Pop()
			xInt, _ := x.(int)
			yInt, _ := y.(int)
			s.Push(yInt - xInt)
		} else if item == '*' {
			x, _ := s.Pop()
			y, _ := s.Pop()
			xInt, _ := x.(int)
			yInt, _ := y.(int)
			s.Push(yInt * xInt)
		} else if item == '/' {
			x, _ := s.Pop()
			y, _ := s.Pop()
			xInt, _ := x.(int)
			yInt, _ := y.(int)
			s.Push(yInt / xInt)
		}
	}
	res, _ := s.Peek()
	result, _ := res.(int)
	return result
}