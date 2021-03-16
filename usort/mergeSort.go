package usort

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
