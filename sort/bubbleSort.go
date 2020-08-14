package sort

func (data *DataArr) BubbleSort() {
	len := len(*data)

	// 需要冒泡的轮数， 每冒泡一轮将最值放到最后边
	// 每处理一轮需要处理的数组规模变小
	for i := len - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if (*data)[j] > (*data)[j+1] {
				(*data)[j], (*data)[j+1] = (*data)[j+1], (*data)[j]
			}
		}
	}
	data.Print()
}

func (data *DataArr) BubbleSort1() {
	len := len(*data)

	//需要冒泡的轮数， 每冒泡一轮将最值放到最后边
	//每处理一轮需要处理的数组规模变小
	for i := 1; i < len; i++ {
		for j := 0; j < len-i; j++ {
			if (*data)[j] > (*data)[j+1] {
				(*data)[j], (*data)[j+1] = (*data)[j+1], (*data)[j]
			}
		}
	}
	data.Print()
}
