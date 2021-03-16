package other

import "github.com/coredumptoday/practice/utils"

func OddEvenSort(arr []int, l, r int) int {
	if arr == nil || len(arr) <= 1 {
		return -1
	}

	odd := l - 1
	idx := 0
	for idx < len(arr) {
		if arr[idx]%2 != 0 {
			utils.IntSliceSwap(arr, odd+1, idx)
			odd++
		}
		idx++
	}
	return odd
}
