package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var s Interface = &Stack{}
	fmt.Println(s.Pop())
	s.Push("yangxuan")
	s.Push("nihao")
	s.Push("nihao2")
	fmt.Println(s)
	if s.Size() != 3 || s.Empty() {
		t.Errorf("Expected size 3 Got %d", s.Size())
	}
	ele, b := s.Pop()
	if ele != "nihao2" || b == false {
		t.Errorf("Expected nihao2 Got %s", ele)
	}
	s.Pop()
	s.Pop()
	fmt.Println(s)
}

func ExampleStack_Push() {
	s := &Stack{}
	s.Push("a")

	// Output:
}

func ExampleStack_Peek() {
	s := &Stack{}
	s.Push("a")
	fmt.Println(s.Peek())
	fmt.Println(s)

	// Output:
	// a true
	// Stack: [a] top
}

func ExampleStack_Pop() {
	s := &Stack{}
	s.Push("a")
	fmt.Println(s.Pop())
	fmt.Println(s)

	// Output:
	// a true
	// Stack: [] top
}