package utils

import (
	"math/rand"
	"time"
)

func GetRandNum() float32 {
	rand.Seed( time.Now().UnixNano() )
	return rand.Float32()
}

func GetValue(maxValue int) int {
	return int( float32(maxValue + 1) * GetRandNum() ) - int( float32(maxValue) * GetRandNum() )
}

func GenerateRandomSlice(maxSize, maxValue int) []int {
	cap := int( float32(maxSize + 1) * GetRandNum() )
	arr := make([]int, cap)

	for i, _ := range arr {
		arr[i] = GetValue(maxValue)
	}

	return arr
}

func IntSliceSwap(arr []int, i, j int)  {
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