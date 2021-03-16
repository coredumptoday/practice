package usort

import "github.com/coredumptoday/practice/xor"

/*
	插入排序
	1. 0 位置上有序
	2. 0 ~ 1 位置上有序
	3. 0 ~ 2 位置上有序
	4. 0 ~ 3 位置上有序
	....
	...
*/
func InsertionSort(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i; j-1 >= 0; j-- {
			if arr[j-1] > arr[j] {
				//utils.IntSliceSwap(arr, j, j - 1)
				xor.Swap(arr, j, j-1)
			}
		}
	}
}
