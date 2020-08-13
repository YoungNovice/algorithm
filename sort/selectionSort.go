package sort

// 选择排序
func (data *DataArr) SelectionSort() {
	len := len(*data)

	for i := 0; i < len - 1; i++ {
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