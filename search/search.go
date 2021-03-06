package search

import "math"

func BinarySortExist(arr []int, v int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}

	if len(arr) == 1 {
		if arr[0] == v {
			return true
		} else {
			return false
		}
	}

	L := 0
	R := len(arr) - 1

	for L < R {
		mid := L + (R-L)/2
		if arr[mid] == v {
			return true
		} else if arr[mid] > v {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}

	return arr[L] == v
}

// 在arr上，找满足>=value的最左位置
func BinarySortSearchLeft(arr []int, value int) int {
	idx := -1
	if arr == nil || len(arr) == 0 {
		return idx
	}

	L := 0
	R := len(arr) - 1

	for L <= R {
		mid := L + (R-L)/2
		if arr[mid] >= value {
			idx = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}

	return idx
}

// 在arr上，找满足<=value的最右位置
func BinarySortSearchRight(arr []int, value int) int {
	idx := -1
	if arr == nil || len(arr) == 0 {
		return idx
	}

	L := 0
	R := len(arr) - 1

	for L <= R {
		mid := L + (R-L)/2
		if arr[mid] <= value {
			idx = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}

	return idx
}

func MaxNumSearch(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return math.MinInt32
	}

	m := process(arr, 0, len(arr)-1)

	return int(m)
}

func process(arr []int, L, R int) float64 {
	if L == R {
		return float64(arr[L])
	}

	mid := L + ((R - L) >> 1)

	left := process(arr, L, mid)
	right := process(arr, mid+1, R)

	return math.Max(left, right)
}
