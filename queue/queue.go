package queue

import "fmt"

type Queue []interface{}

func (q* Queue) Enqueue(a interface{}) {
	*q = append(*q, a)
}

func (q* Queue) Dequeue() (interface{}, bool) {
	front, b := q.Front()
	if b {
		*q = (*q)[1:]
	}
	return front, b
}

func (q Queue) Front() (interface{}, bool) {
	if !q.Empty() {
		return q[0], true
	}
	return nil, false
}

func (q Queue) Size() int {
	return len(q)
}

func (q Queue) Empty() bool {
	return len(q) == 0
}

func (q Queue) String() string {
	return fmt.Sprintf("Queue: front %v end",  []interface{}(q))
}





