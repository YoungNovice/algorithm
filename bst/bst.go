package bst

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

	if e.CompareTo(n.e) < 0 {
		n.left = b.add(n.left, e)
	} else if e.CompareTo(n.e) > 0 {
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
	} else if e.CompareTo(n.e) < 0 {
		return b.contains(n.left, e)
	} else {
		return b.contains(n.right, e)
	}
}
