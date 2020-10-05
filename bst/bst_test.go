package bst

import (
	"testing"
)

type myInt int

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



