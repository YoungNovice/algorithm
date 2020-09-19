package stack

import "fmt"

type Stack []interface{}

// 实现stack_interface
func (s* Stack) Push(a interface{}) {
	*s = append(*s, a)
}

func (s* Stack) Pop() (interface{}, bool) {
	ele, b := s.Peek()
	if b {
		*s = (*s)[:s.Size()-1]
	}
	return ele, b
}

func (s Stack) Peek() (interface{}, bool) {
	if !s.Empty() {
		lastEle := len(s) - 1
		last := s[lastEle]
		return last, true
	}
	return nil, false
}

func (s Stack) Size() int {
	return len(s)
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

// 实现Stringer中 String方法
func (s Stack) String() string {
	return fmt.Sprintf("Stack: %v top", []interface{}(s))
}
