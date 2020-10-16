package set

import (
	"algorithm/bst"
	"algorithm/common"
)

type Interface interface {
	Add(e common.Comparable)
	Remove(e common.Comparable)
	Contains(e common.Comparable) bool
	Size() int
	Empty() bool
}

type BSTSet struct {
	bst bst.BST
}

func (bs *BSTSet) Add(e common.Comparable) {
	bs.bst.Add(e)
}

func (bs *BSTSet) Remove(e common.Comparable) {
	bs.bst.Remove(e)
}

func (bs *BSTSet) Contains(e common.Comparable) bool {
	return bs.bst.Contains(e)
}

func (bs *BSTSet) Size() int {
	return bs.bst.Size()
}

func (bs *BSTSet) Empty() bool {
	return bs.Size() == 0
}

