package usort

import "github.com/coredumptoday/practice/xor"

/*
	冒泡排序
	1. 从 0 ~ N-1 位置中选择一个最大的放在 N-1 位置
	2. 从 0 ~ N-2 位置中选择一个最大的放在 N-2 位置
	3. 从 0 ~ N-3 位置中选择一个最大的放在 N-3 位置
	...
	...
*/
func BubbleSort(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}

	for e := len(arr) - 1; e > 0; e-- {
		tMax := e
		for j := e - 1; j >= 0; j-- {
			if arr[j] > arr[tMax] {
				tMax = j
			}
		}
		//utils.IntSliceSwap(arr, e, tMax)
		xor.Swap(arr, e, tMax)
	}
}
