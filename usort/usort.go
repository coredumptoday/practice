package usort

import (
	"github.com/coredumptoday/practice/heap"
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

func MergeSort(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	processMergeSort(arr, 0, len(arr)-1)
}

func processMergeSort(arr []int, L, R int) {
	if L == R {
		return
	}
	mid := L + (R-L)/2
	processMergeSort(arr, L, mid)
	processMergeSort(arr, mid+1, R)
	merge(arr, L, mid, R)
}

func merge(arr []int, L, M, R int) {
	helpArr := make([]int, R-L+1)
	hIdx := 0
	p1 := L
	p2 := M + 1

	for p1 <= M && p2 <= R {
		if arr[p1] <= arr[p2] {
			helpArr[hIdx] = arr[p1]
			p1++
		} else {
			helpArr[hIdx] = arr[p2]
			p2++
		}
		hIdx++
	}

	for p1 <= M {
		helpArr[hIdx] = arr[p1]
		p1++
		hIdx++
	}

	for p2 <= R {
		helpArr[hIdx] = arr[p2]
		p2++
		hIdx++
	}

	for i, v := range helpArr {
		arr[L+i] = v
	}
}

func HeapSort(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}

	h := heap.BuildHeapFromSlice(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		v, _ := h.Pop()
		arr[i] = v
	}
}
