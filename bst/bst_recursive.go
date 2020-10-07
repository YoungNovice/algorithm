package bst


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
