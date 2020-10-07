package bst

// 先序遍历非递归实现
// 非递归1
// (先放右孩子 再放左孩子) 出栈时机是左孩子先出栈 所以先右后左
// 同时由于先序遍历是先访问根的所以看着比较简单
func (b *BST) PreOrderNR1(f func(*node)) {
	var s []*node
	s = append(s, b.root)

	for len(s) > 0 {
		// pop
		top := s[len(s)-1]
		f(top)
		s = s[:len(s)-1]
		// push right if exists
		if top.right != nil {
			s = append(s, top.right)
		}
		// push left if exists
		if top.left != nil {
			s = append(s, top.left)
		}
	}
}

func (b *BST) PreOrderNR2(f func(*node)) {
	var s []*node
	s = append(s, b.root)

	for len(s) > 0 {
		// pop
		top := s[len(s)-1]
		f(top)
		s = s[:len(s)-1]
		// push right if exists
		if top.right != nil {
			s = append(s, top.right)
		}
		// push left if exists
		if top.left != nil {
			s = append(s, top.left)
		}
	}
}

// 中序遍历非递归 一直向左遍历到底 底部出栈 出栈后访问结点，之后转到右子树
func (b *BST) InOrderNR(f func(*node)) {
	if b.root == nil {
		return
	}

	var s []*node
	cur := b.root
	for len(s) > 0 || cur != nil {
		// 当前结点有左孩子一路push
		for cur != nil {
			s = append(s, cur)
			cur = cur.left
		}
		// end to left bottom
		if len(s) > 0 {
			// pop
			cur = s[len(s)-1]
			s = s[:len(s)-1]
			f(cur)
			cur = cur.right
			// if cur is nil go to next loop pop again
		}
	}
}

func (b *BST) InOrderNR1(f func(*node)) {
	if b.root == nil {
		return
	}
	var s []*node
	cur := b.root

	for len(s) > 0 || cur != nil {
		if cur != nil {
			s = append(s, cur)
			cur = cur.left
		} else {
			// pop
			cur = s[len(s)-1]
			s = s[:len(s)-1]
			f(cur)
			cur = cur.right
		}
	}
}
