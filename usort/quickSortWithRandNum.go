package usort

import "github.com/coredumptoday/practice/utils"

func processWithRandNum(arr []int, l, r int) {
	if l >= r {
		return
	}
	x := int(utils.GetRandNum() * float32(r-l+1))
	utils.IntSliceSwap(arr, l+x, r)
	m := split3range(arr, l, r)
	processWithRandNum(arr, l, m[0]-1)
	processWithRandNum(arr, m[1]+1, r)
}

func QuickSortWithRandNum(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	processWithRandNum(arr, 0, len(arr)-1)
}
