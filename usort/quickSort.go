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

func process1(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := split2range(arr, l, r)
	process1(arr, l, m-1)
	process1(arr, m+1, r)
}

func QuickSort1(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	process1(arr, 0, len(arr)-1)
}

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

func process2(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := split3range(arr, l, r)

	process2(arr, l, m[0]-1)
	process2(arr, m[1]+1, r)
}

func QuickSort2(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	process2(arr, 0, len(arr)-1)
}

func process3(arr []int, l, r int) {
	if l >= r {
		return
	}
	x := int(utils.GetRandNum() * float32(r-l+1))
	utils.IntSliceSwap(arr, l+x, r)
	m := split3range(arr, l, r)
	process3(arr, l, m[0]-1)
	process3(arr, m[1]+1, r)
}

func QuickSort3(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	process3(arr, 0, len(arr)-1)
}
