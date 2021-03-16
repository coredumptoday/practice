package usort

import "github.com/coredumptoday/practice/utils"

func split3range(arr []int, l, r int) []int {
	if l > r {
		return []int{-1, -1}
	}
	if l == r {
		return []int{l, r}
	}

	less := l - 1
	more := r
	idx := l
	for idx < more {
		if arr[idx] < arr[r] {
			utils.IntSliceSwap(arr, less+1, idx)
			less++
			idx++
		} else if arr[idx] > arr[r] {
			utils.IntSliceSwap(arr, more-1, idx)
			more--
		} else {
			idx++
		}
	}

	utils.IntSliceSwap(arr, more, r)
	return []int{less + 1, more}
}

func processWith3Range(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := split3range(arr, l, r)

	processWith3Range(arr, l, m[0]-1)
	processWith3Range(arr, m[1]+1, r)
}

func QuickSortWith3Range(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	processWith3Range(arr, 0, len(arr)-1)
}
