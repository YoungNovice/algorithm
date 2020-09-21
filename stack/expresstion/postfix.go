package expresstion

import "algorithm/stack"

// 后缀表达式求值最简单
// 遇到数字入栈 遇到运算符 从栈中区两个操作数 先出来的放后边 运算结果入栈
// 最后栈里面的数字就是表达式的结果
func PostfixEval(str string) int {
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