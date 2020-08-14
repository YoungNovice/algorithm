package sort

// 选择排序
func (data *DataArr) SelectionSort() {
	len := len(*data)

	// 每轮寻找最小值的下标 将最小值交换到最前面 接着处理后边的
	for i := 0; i < len-1; i++ {
		minIndex := i

		for j := i + 1; j < len; j++ {
			if (*data)[j] < (*data)[minIndex] {
				minIndex = j
			}
		}

		(*data)[minIndex], (*data)[i] = (*data)[i], (*data)[minIndex]
	}

	data.Print()
}
