package stack

type Interface interface {
	Push(a interface{})
	Pop() (interface{}, bool)
	Peek() (interface{}, bool)
	Size() int
	Empty() bool
}
