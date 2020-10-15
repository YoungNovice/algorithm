package bst

import (
	"fmt"
	"strings"
)

type Comparable interface {
	CompareTo(c Comparable) int
}

type node struct {
	e Comparable
	left, right *node
}

type BST struct {
	root *node
	size int
}

// 向二分搜索树中添加元素
func (b *BST) Add(e Comparable) {
	b.root = b.add(b.root, e)
}

// 递归添加 这个比较难理解 对于相同的值丢弃
func (b *BST) add(n *node, e Comparable) *node {
	if n == nil {
		b.size++
		return &node{e, nil, nil}
	}

	if n.e.CompareTo(e) > 0 {
		n.left = b.add(n.left, e)
	} else if n.e.CompareTo(e) < 0 {
		n.right = b.add(n.right, e)
	}
	return n
}

func (b *BST) Contains(e Comparable) bool {
	return b.contains(b.root, e)
}

func (b *BST) contains(n *node, e Comparable) bool {
	if n == nil {
		return false
	}

	if e.CompareTo(n.e) == 0 {
		return true
	} else if n.e.CompareTo(e) > 0 {
		return b.contains(n.left, e)
	} else {
		return b.contains(n.right, e)
	}
}

// 层次遍历 利用队列
func (b *BST) LevelOrder(f func(*node)) {
	var q []*node
	q = append(q, b.root)
	for len(q) > 0 {
		// dequeue
		cur := q[0]
		q = q[1:]
		f(cur)
		// enqueue
		if cur.left != nil {
			q = append(q, cur.left)
		}
		if cur.right != nil {
			q = append(q, cur.right)
		}
	}
}

func (b *BST) Min() *node {
	if b.size == 0 {
		panic("BST is empty!")
	}
	return b.min(b.root)
}

func (b *BST) min(n *node) *node {
	if n.left == nil {
		return n
	}
	return b.min(n.left)
}

func (b *BST) Max() *node {
	if b.size == 0 {
		panic("BST is empty!")
	}
	return b.max(b.root)
}

func (b *BST) max(n *node) *node {
	if n.right == nil {
		return n
	}
	return b.max(n.right)
}

func (b BST) String() string {
	var sb strings.Builder
	generateBSTString(b.root, 0, &sb)
	return sb.String()
}

func generateBSTString(n *node, depth int, sb *strings.Builder) {
	if n == nil {
		sb.WriteString(strings.Repeat("--", depth) + "null\n")
		return
	}
	sb.WriteString(fmt.Sprintf("%s%s%s", strings.Repeat("--", depth), n.e, "\n"))
	generateBSTString(n.left, depth+1, sb)
	generateBSTString(n.right, depth+1, sb)
}

