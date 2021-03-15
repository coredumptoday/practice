package other

/**
在一个数组中，一个数左边比它小的数的总和，叫数的小和，所有数的小和累加起来，叫数组小和。求数组小和。
例子： [1,3,4,2,5]
1左边比1小的数：没有
3左边比3小的数：1
4左边比4小的数：1、3
2左边比2小的数：1
5左边比5小的数：1、3、4、 2
所以数组的小和为1+1+3+1+1+3+4+2=16
*/
func LessSum(arr []int) int {
	if arr == nil || len(arr) <= 1 {
		return 0
	}
	return processLessSum(arr, 0, len(arr)-1)
}

func processLessSum(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	mid := L + (R-L)/2
	return processLessSum(arr, L, mid) +
		processLessSum(arr, mid+1, R) +
		merge(arr, L, mid, R)
}

func merge(arr []int, L, M, R int) int {
	helpArr := make([]int, R-L+1)
	hIdx := 0
	p1 := L
	p2 := M + 1
	res := 0

	for p1 <= M && p2 <= R {
		if arr[p1] < arr[p2] {
			res += (R - p2 + 1) * arr[p1]
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

	return res
}

func Comparator(arr []int) int {
	res := 0
	for i, v := range arr {
		if i == 0 {
			continue
		}
		for j := 0; j < i; j++ {
			if arr[j] < v {
				res += arr[j]
			}
		}
	}
	return res
}
