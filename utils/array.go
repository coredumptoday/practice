package utils

import "math"

func GenerateRandomSlice(maxSize, maxValue int) []int {
	cap := int(float32(maxSize+1) * GetRandNum())
	arr := make([]int, cap)

	for i, _ := range arr {
		arr[i] = GetValue(maxValue)
	}

	return arr
}

func IntSliceSwap(arr []int, i, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}

func IntInSlice(arr []int, v int) bool {
	if arr == nil || len(arr) <= 0 {
		return false
	}

	for _, val := range arr {
		if val == v {
			return true
		}
	}

	return false
}

func IsSliceEqual(a1, a2 []int) bool {
	if (a1 == nil && a2 != nil) || (a1 != nil && a2 == nil) {
		return false
	}

	if a1 == nil && a2 == nil {
		return true
	}

	if len(a1) != len(a2) {
		return false
	}

	for i, v1 := range a1 {
		if a2[i] != v1 {
			return false
		}
	}

	return true
}

func mkRandomNum(a int) int {
	return int((float32(a)*GetRandNum() + 1) - (float32(a)*GetRandNum() + 1))
}

func MKRandomSlice(maxKinds, a, k, m int) ([]int32, int32) {
	ktimeNum := mkRandomNum(a)

	// 真命天子出现的次数
	times := 0
	if GetRandNum() < 0.5 {
		times = k
	} else {
		times = int(GetRandNum()*float32(m-1) + 1)
	}

	// 2
	numKinds := int(GetRandNum()*float32(maxKinds) + 2)

	// k * 1 + (numKinds - 1) * m
	arr := make([]int32, times+(numKinds-1)*m)

	index := 0
	for ; index < times; index++ {
		arr[index] = int32(ktimeNum)
	}

	set := make(map[int]bool, numKinds)
	set[ktimeNum] = true
	numKinds--

	for numKinds != 0 {
		curNum := 0
		for {
			curNum = mkRandomNum(a)
			if !set[curNum] {
				break
			}
		}
		set[curNum] = true
		numKinds--
		for i := 0; i < m; i++ {
			arr[index] = int32(curNum)
			index++
		}
	}

	for i := 0; i < len(arr); i++ {
		j := int(GetRandNum() * float32(len(arr)))
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	}

	return arr, int32(ktimeNum)
}

func KmTest(arr []int32, k int32) int32 {
	mm := make(map[int32]int32, len(arr))

	for _, v := range arr {
		if n, ok := mm[v]; ok {
			mm[v] = n + 1
		} else {
			mm[v] = 1
		}
	}

	for i, v := range mm {
		if v == k {
			return i
		}
	}

	return math.MinInt32
}
