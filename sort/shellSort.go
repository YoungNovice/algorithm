package sort

// 希尔排序 取间隔排序
func (data *DataArr) ShellSort() {
	len := len(*data)

	// Knuth序列
	h := 1
	for h <= len/3 {
		h = h*3 + 1
	}
	for gap := h; gap > 0; gap = (gap - 1) / 3 {
		for i := gap; i < len; i++ {
			for j := i; j > gap-1; j -= gap {
				if (*data)[j] < (*data)[j-1] {
					(*data)[j], (*data)[j-1] = (*data)[j-1], (*data)[j]
				}
			}
		}
	}
	data.Print()
}
