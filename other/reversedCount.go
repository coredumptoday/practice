package other

func ReversedCount(arr []int) int {
	if arr == nil || len(arr) <= 1 {
		return 0
	}
	return processReversedCount(arr, 0, len(arr)-1)
}

func processReversedCount(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	mid := L + (R-L)/2
	return processReversedCount(arr, L, mid) +
		processReversedCount(arr, mid+1, R) +
		revMerge(arr, L, mid, R)
}

func revMerge(arr []int, L, M, R int) int {
	helpArr := make([]int, R-L+1)
	hIdx := len(helpArr) - 1
	p1 := M
	p2 := R
	res := 0

	for p1 >= L && p2 > M {
		if arr[p1] > arr[p2] {
			res += p2 - M
			helpArr[hIdx] = arr[p1]
			p1--
		} else {
			helpArr[hIdx] = arr[p2]
			p2--
		}
		hIdx--
	}

	for p1 >= L {
		helpArr[hIdx] = arr[p1]
		p1--
		hIdx--
	}

	for p2 > M {
		helpArr[hIdx] = arr[p2]
		p2--
		hIdx--
	}

	for i, v := range helpArr {
		arr[L+i] = v
	}

	return res
}

func RevComparator(arr []int) int {
	res := 0
	for i, v := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < v {
				res++
			}
		}
	}
	return res
}
