package usort

import (
	"github.com/coredumptoday/practice/utils"
)

func split2range(arr []int, l, r int) int {
	if l > r {
		return -1
	}
	if l == r {
		return r
	}

	lessEqual := l - 1
	idx := l
	for idx < r {
		if arr[idx] <= arr[r] {
			utils.IntSliceSwap(arr, lessEqual+1, idx)
			lessEqual++
		}
		idx++
	}

	utils.IntSliceSwap(arr, lessEqual+1, r)
	lessEqual++
	return lessEqual
}

func processWith2Range(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := split2range(arr, l, r)
	processWith2Range(arr, l, m-1)
	processWith2Range(arr, m+1, r)
}

func QuickSortWith2Range(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	processWith2Range(arr, 0, len(arr)-1)
}
