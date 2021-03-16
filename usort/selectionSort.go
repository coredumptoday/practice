package usort

import (
	"github.com/coredumptoday/practice/xor"
)

/*
	选择排序
	1. 从 0 ~ N-1 位置选择一个最小的放在 0 位置
	2. 从 1 ~ N-1 位置选择一个最小的放在 1 位置
	3. 从 2 ~ N-1 位置选择一个最小的放在 2 位置
	...
	...
*/
func SelectionSort(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		tMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[tMin] {
				tMin = j
			}
		}
		//utils.IntSliceSwap(arr, i, tMin)
		xor.Swap(arr, i, tMin)
	}
}
