package sort

func (data *DataArr) InsertionSort()  {
	len := len(*data)

	// 取一个数向前边的子数组中插入
	// 插完之后前边的数是有序状态， 跟冒泡差不多
	// 不同的是冒泡是做相邻比较每轮得出一个最值 插入是向有序的数组中插入 插完还是有序的
	// 只不过是有数组存储的特性 插入的位置需要依次挪动 所以看起来像冒泡
	for i := 1; i < len; i++ {
		for j := i; j > 0; j-- {
			if (*data)[j] < (*data)[j-1] {
				(*data)[j], (*data)[j-1] = (*data)[j-1], (*data)[j]
			}
		}
	}
	data.Print()
}
