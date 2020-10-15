package bst

// 先序遍历非递归实现
// 非递归1
// (先放右孩子 再放左孩子) 出栈时机是左孩子先出栈 所以先右后左
// 同时由于先序遍历是先访问根的所以看着比较简单
func (b *BST) PreOrderNR0(f func(*node)) {
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

func (b *BST) PreOrderNR1(f func(*node)) {
	if b.root == nil {
		return
	}

	var s []*node
	cur := b.root
	for cur != nil || len(s) > 0 {
		for cur != nil {
			f(cur)
			// push
			s = append(s, cur)
			cur = cur.left
		}
		if len(s) > 0 {
			// pop
			cur = s[len(s)-1]
			s = s[:len(s)-1]
			cur = cur.right
		}
	}
}

// 中序遍历非递归 一直向左遍历到底 底部出栈 出栈后访问结点，之后转到右子树
func (b *BST) InOrderNR1(f func(*node)) {
	if b.root == nil {
		return
	}

	var s []*node
	cur := b.root
	for cur != nil || len(s) > 0 {
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

// 后续遍历比较麻烦 需要一个结点需要入两次栈
func (b *BST) PostOrderNR1(f func(*node)) {
	if b.root == nil {
		return
	}
	type PostNode struct {
		n    *node
		flag int
	}
	var s []*PostNode
	cur := b.root

	for len(s) > 0 || cur != nil {
		if cur != nil {
			// 第一次push
			s = append(s, &PostNode{cur, 1})
			cur = cur.left
		} else {
			// pop
			top := s[len(s)-1]
			s = s[:len(s)-1]
			cur = top.n
			if top.flag == 1 {
				top.flag = 2
				// 第二次push
				s = append(s, top)
				// 转向右子树
				cur = cur.right
			} else {
				// flag == 2
				f(cur)
				cur = nil
			}
		}
	}
}

// 后续遍历比较麻烦 需要一个结点需要入两次栈
func (b *BST) PostOrderNR2(f func(*node)) {
	if b.root == nil {
		return
	}
	type PostNode struct {
		n    *node
		flag int
	}
	var s []*PostNode
	cur := b.root

	for len(s) > 0 || cur != nil {
		for cur != nil {
			s = append(s, &PostNode{cur, 1})
			cur = cur.left
		}
		// pop
		top := s[len(s)-1]
		s = s[:len(s)-1]
		cur = top.n

		if top.flag == 1 {
			// 第一次push
			top.flag = 2
			// 第二次push
			s = append(s, top)
			// 转向右子树
			cur = cur.right
		} else {
			f(cur)
			cur = nil
		}
	}
}
