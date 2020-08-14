package sort

// 希尔排序 取间隔排序
func (data *DataArr) ShellSort()  {
	len := len(*data)
	for gap := len >> 1; gap > 0; gap /= 2 {
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
