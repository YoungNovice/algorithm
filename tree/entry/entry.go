package main

import (
	"algorithm/tree"
	"algorithm/tree/wrapper"
	"fmt"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	// 遍历
	root.Traverse()

	// 计数
	nodeCount := 0
	root.TraverseFunc(func(n *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount :", nodeCount)

	// chan
	c := root.TraverseWithChan()
	max := 0
	for n := range c {
		if n.Value > max {
			max = n.Value
		}
	}
	fmt.Printf("max Node is %d\n", max)

	myTreeNode := wrapper.MyTreeNode{Node: &root}
	myTreeNode.PostOrder()
	fmt.Println()
}
