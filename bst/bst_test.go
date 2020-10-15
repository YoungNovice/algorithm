package bst

import (
	"fmt"
	"strconv"
	"testing"
)

type myInt int

func (m myInt) String() string {
	return strconv.Itoa(int(m))
}

func getData() BST {
	var bst BST
	bst.Add(myInt(5))
	bst.Add(myInt(3))
	bst.Add(myInt(6))
	bst.Add(myInt(8))
	bst.Add(myInt(4))
	bst.Add(myInt(2))
	return bst
}

func (m myInt) CompareTo(c Comparable) int {
	if val, ok := c.(myInt); ok {
		if m > val {
			return 1
		} else if m < val {
			return -1
		} else {
			return 0
		}
	} else {
		panic("parameter error")
	}
}

func TestBST_Add(t *testing.T) {
	var bst BST
	bst.Add(myInt(50))
	bst.Add(myInt(20))
	bst.Add(myInt(30))
	bst.Add(myInt(25))
}

func TestBST_PreOrder(t *testing.T) {
	bst := getData()
	// 递归
	bst.PreOrder(func(n *node) {
		fmt.Print(n.e, " ")
	})
	fmt.Println()
	// 非递归1
	bst.PreOrderNR0(func(n *node) {
		fmt.Print(n.e, " ")
	})
	fmt.Println()
	// 非递归2
	bst.PreOrderNR1(func(n *node) {
		fmt.Print(n.e, " ")
	})
	fmt.Println()
}



func TestBST_InOrder(t *testing.T) {
	bst := getData()
	// 递归
	bst.InOrder(func(n *node) {
		fmt.Print(n.e, " ")
	})
	fmt.Println()
	// 非递归
	bst.InOrderNR1(func(n *node) {
		fmt.Print(n.e, " ")
	})
	fmt.Println()

}

func TestBST_PostOrder(t *testing.T) {
	bst := getData()
	// 递归
	bst.PostOrder(func(n *node) {
		fmt.Print(n.e, " ")
	})
	fmt.Println()
	// 非递归
	bst.PostOrderNR1(func(n *node) {
		fmt.Print(n.e, " ")
	})
	fmt.Println()
}

func TestBST_LevelOrder(t *testing.T) {
	bst := getData()
	bst.LevelOrder(func(n *node) {
		fmt.Print(n.e, " ")
	})
	fmt.Println()
}

func TestBST_Min(t *testing.T) {
	bst := getData()
	bst.InOrderNR1(func(n *node) {
		fmt.Print(n.e, " ")
	})
	fmt.Println()
	fmt.Println(bst.Min().e, bst.MinNR().e)
	fmt.Println(bst.Max().e, bst.MaxNR().e)
}



