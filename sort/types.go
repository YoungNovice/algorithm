package sort

import "fmt"

type DataArr []int

func (data *DataArr) Print()  {
	if data == nil {
		return
	}

	for _, item := range *data {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

