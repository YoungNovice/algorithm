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

// 先序遍历
func (b *BST) PreOrder(f func(*node)) {
	b.preOrder(b.root, f)
}

func (b *BST) preOrder(n *node, f func(*node)) {
	if n == nil {
		return
	}
	f(n)
	b.preOrder(n.left, f)
	b.preOrder(n.right, f)
}

// 中序遍历
func (b *BST) InOrder(f func(*node)) {
	b.inOrder(b.root, f)
}

func (b *BST) inOrder(n *node, f func(*node)) {
	if n == nil {
		return
	}
	b.inOrder(n.left, f)
	f(n)
	b.inOrder(n.right, f)
}

// 后序遍历
func (b *BST) PostOrder(f func(*node)) {
	b.postOrder(b.root, f)
}

func (b *BST) postOrder(n *node, f func(*node)) {
	if n == nil {
		return
	}
	b.postOrder(n.left, f)
	b.postOrder(n.right, f)
	f(n)
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

