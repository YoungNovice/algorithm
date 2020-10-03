package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList_addFirst(t *testing.T) {
	var l LinkedList
	l.addFirst("1")
	l.addFirst("2")
	l.addLast("last")
	l.set(1, "haha")
	fmt.Println(l, l.getFirst(), l.getLast())
}
