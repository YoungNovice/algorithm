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

// 删除二分搜索树中的结点
// Hibbard Deleion
// 用右子树的最小值替代原位置 或者左子树的最大值
func (b *BST) Remove(e Comparable) {
	b.root = b.remove(b.root, e)
}

func (b *BST) remove(n *node, e Comparable) *node {
	if n.e.CompareTo(e) > 0 {
		n.left = b.remove(n.left, e)
		return n
	} else if n.e.CompareTo(e) < 0 {
		n.right = b.remove(n.right, e)
		return n
	} else {
		if n.left == nil {
			right := n.right
			n.right = nil
			b.size--
			return right
		}
		if n.right == nil {
			left := n.left
			n.left = nil
			b.size--
			return left
		}
		// 取右边的最小 or 取左边最大
		//min := b.min(n.right)
		//min.right = b.removeMin(n.right)
		//min.left = n.left
		//return min
		max := b.max(n.left)
		max.left = b.removeMax(n.left)
		max.right = n.right
		return max
	}
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

// 删除最小值
func (b *BST) RemoveMin() *node {
	e := b.Min()
	b.root = b.removeMin(b.root)
	return e
}

func (b *BST) removeMin(n *node) *node {
	if n.left == nil {
		right := n.right
		n.right = nil
		b.size--
		return right
	}
	n.left = b.removeMin(n.left)
	return n
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

// 删除最大值
func (b *BST) RemoveMax() *node {
	e := b.Max()
	b.root = b.removeMax(b.root)
	return e
}

func (b *BST) removeMax(n *node) *node {
	if n.right == nil {
		left := n.left
		n.left = nil
		b.size--
		return left
	}
	n.right = b.removeMax(n.right)
	return n
}

// 查询最大值
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

