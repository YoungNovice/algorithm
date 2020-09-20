package wrapper

import "algorithm/tree"

// 扩展 tree.Node结构
// 加一个后序遍历
type MyTreeNode struct {
	Node *tree.Node
}

func (myTreeNode *MyTreeNode) PostOrder() {
	if myTreeNode == nil || myTreeNode.Node == nil {
		return
	}
	myTreeNodeLeft := &MyTreeNode{myTreeNode.Node.Left}
	myTreeNodeLeft.PostOrder()
	myTreeNodeRight := &MyTreeNode{myTreeNode.Node.Right}
	myTreeNodeRight.PostOrder()
	myTreeNode.Node.Print()
}
