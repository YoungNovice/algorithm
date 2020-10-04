package linkedlist

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	dummyHead *node
	size int
}

type node struct {
	e interface{}
	next *node
}

func (l *LinkedList) opCheck() {
	if l.dummyHead == nil {
		l.dummyHead = &node{nil, nil}
		l.size = 0
	}
}

func (l LinkedList) String() string {
	l.opCheck()
	var b strings.Builder
	cur := l.dummyHead.next
	for cur != nil {
		b.WriteString(fmt.Sprint(cur.e, "->"))
		cur = cur.next
	}
	b.WriteString("NULL")
	return b.String()
}

func (l *LinkedList) add(index int, e interface{}) {
	l.opCheck()
	if index < 0 || index > l.size {
		panic("valid index")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = &node{e, prev.next}
	l.size++
}

func (l *LinkedList) Delete(index int) {
	l.opCheck()
	// 增加的判断 和更新删除不一样 增加是可以在最后一个位置加的
	if index < 0 || index > l.size-1 {
		panic("valid index")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	l.size--
}

func (l *LinkedList) AddFirst(e interface{}) {
	l.add(0, e)
}

func (l *LinkedList) AddLast(e interface{}) {
	l.add(l.size, e)
}

func (l *LinkedList) set(index int, e interface{}) {
	cur := l.get(index)
	cur.e = e
}

func (l *LinkedList) Empty() bool {
	return l.size == 0
}

func (l *LinkedList) get(index int) *node {
	if index < 0 || index > l.size-1 {
		panic("valid index")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (l *LinkedList) First() interface{} {
	return l.get(0).e
}

func (l *LinkedList) Last() interface{} {
	return l.get(l.size-1).e
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next
	for cur != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false
}
