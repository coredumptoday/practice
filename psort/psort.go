package psort

import "github.com/coredumptoday/practice/utils"

/*
	选择排序
	1. 从 0 ~ N-1 位置选择一个最小的放在 0 位置
	2. 从 1 ~ N-1 位置选择一个最小的放在 1 位置
	3. 从 2 ~ N-1 位置选择一个最小的放在 2 位置
	...
	...
*/
func SelectionSort(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		tMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[tMin] {
				tMin = j
			}
		}
		utils.IntSliceSwap(arr, i, tMin)
	}

	/*for i := len(arr) - 1; i >= 0; i-- {
		tMax := i
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[tMax] {
				tMax = j
			}
		}
		utils.IntSliceSwap(arr, i, tMax)
	}*/
}

/*
	冒泡排序
	1. 从 0 ~ N-1 位置中选择一个最大的放在 N-1 位置
	2. 从 0 ~ N-2 位置中选择一个最大的放在 N-2 位置
	3. 从 0 ~ N-3 位置中选择一个最大的放在 N-3 位置
	...
	...
*/
func BubbleSort(arr []int) {
	for e := len(arr) - 1; e > 0; e-- {
		tMax := e
		for j := e - 1; j >= 0; j-- {
			if arr[j] > arr[tMax] {
				tMax = j
			}
		}
		utils.IntSliceSwap(arr, e, tMax)
	}
}