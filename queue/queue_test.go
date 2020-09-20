package queue

import (
	"fmt"
	"testing"
)

func TestSliceQueue(t *testing.T) {
	var q Interface = &Queue{}
	q.Enqueue("1")
	q.Enqueue("2")
	q.Enqueue(3)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
}

func BenchmarkSliceQueue(b *testing.B) {
	q := &Queue{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < b.N; i++ {
		if _, b := q.Dequeue(); b == false {
			fmt.Println("no val in queue")
		}
	}
}
