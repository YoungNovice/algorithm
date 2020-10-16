package set

import (
	"algorithm/common"
	"fmt"
	"testing"
)

type myInt int

func (m myInt) CompareTo(c common.Comparable) int {
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

func TestBSTSet_Add(t *testing.T) {
	var bs BSTSet
	bs.Add(myInt(10))
	bs.Add(myInt(20))
	fmt.Println(bs.Size())
}
