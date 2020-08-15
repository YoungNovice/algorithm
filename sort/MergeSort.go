package sort

// 归并排序
func (data *DataArr) MergeSort()  {
	len := len(*data)
	sort(data, 0, len-1)
}


func sort(data *DataArr, start, end int) {
	if end == start {
		return
	}
	mid := start + (end- start) / 2
	// 先排左边的
	sort(data, start, mid)
	// 再排右边的
	sort(data, mid+1, end)
	// 左右归并
	merge(data, start, mid+1, end)
}

func merge(data *DataArr, left, right, end int)  {
	mid := right-1
	subSize := end - left + 1
	subData := make(DataArr, subSize)
	l := left
	r := right
	k := 0
	// 将左右数据归并到subData中
	for l <= mid && r <= end {
		if (*data)[l] <= (*data)[r] {
			subData[k] = (*data)[l]
			k++
			l++
		} else {
			subData[k] = (*data)[r]
			k++
			r++
		}
	}

	for l <= mid {
		subData[k] = (*data)[l]
		k++
		l++
	}

	for r <= end {
		subData[k] = (*data)[r]
		k++
		r++
	}

	for i := 0; i < subSize; i++ {
		(*data)[left+i] = subData[i]
	}

}
