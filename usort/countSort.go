package usort

func CountSort(arr []int, max int) {
	bucket := make([]int, max+1)
	for _, v := range arr {
		bucket[v]++
	}

	idx := 0
	for ii, v := range bucket {
		for i := 1; i <= v; i++ {
			arr[idx] = ii
			idx++
		}
	}
}
