package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList_addFirst(t *testing.T) {
	var l LinkedList
	l.AddFirst("1")
	l.AddFirst("2")
	l.AddLast("last")
	l.set(1, "haha")
	fmt.Println(l, l.First(), l.Last())
	l.Delete(0)
	fmt.Println(l, l.First(), l.Last())
}
