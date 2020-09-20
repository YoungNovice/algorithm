package queue

type Interface interface {
	Enqueue(interface{})
	Dequeue() (interface{}, bool)
	Front() (interface{}, bool)
	Size() int
	Empty() bool
}
