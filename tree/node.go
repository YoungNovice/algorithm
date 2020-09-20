package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) Traverse() {
	// 采用中序遍历
	// node.traverseMid()
	fmt.Println("---------")
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
	fmt.Println("---------")
}

// 遍历
func (node *Node) TraverseFunc(f func(node *Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChan() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(n *Node) {
			out <- n
		})
		close(out)
	}()
	return out
}

// 中序遍历
func (node *Node) traverseMid() {
	if node == nil {
		return
	}
	node.Left.traverseMid()
	node.Print()
	node.Right.traverseMid()
}

// 没有构造函数用工厂方法替代
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
