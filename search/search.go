package search

func BinarySortExist(arr []int, v int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}

	if len(arr) == 1 {
		if arr[0] == v {
			return true
		}else {
			return false
		}
	}

	L := 0
	R := len(arr) - 1

	for L < R {
		mid := L + (R - L) / 2
		if arr[mid] == v {
			return true
		}else if arr[mid] > v {
			R = mid - 1
		}else {
			L = mid + 1
		}
	}

	return arr[L] == v
}