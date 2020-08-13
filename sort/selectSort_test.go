package sort

import "testing"

func TestDataArr_SelectionSort(t *testing.T) {
	var s DataArr = []int{4, 5, 100, 3, 20, 40}
	s.SelectionSort()
	s.Print()
	t.Log("ok")
}
