package xor

import (
	"math"
)

func Swap(arr []int, i, j int) {
	if arr[i] == arr[j] { //相同的数字异或结果为0，该条件要独立判断
		return
	}
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

//获取唯一出现奇数次的数字
func SearchOnlyOneShowOddNumFromSlice(arr []int) int {
	num := 0
	for _, v := range arr {
		num = num ^ v
	}
	return num
}

//获取两个出现奇数次的数字
func SearchTwoShowOddNumFromSlice(arr []int) (int, int) {
	xor := 0
	for _, v := range arr {
		xor = xor ^ v
	}

	rightOne := xor & (-xor)	//取出最右的1，也就是最小的不同位
	x := 0
	for _, v := range arr {
		if v & rightOne != 0 {	//根据最小不同位将数组分位两组
			x = x ^ v
		}
	}

	return x, xor ^ x
}

/**
	一个数组中有一种数出现K次，其他数都出现了M次，
	M > 1,  K < M
	找到，出现了K次的数，
	要求，额外空间复杂度O(1)，时间复杂度O(N)
 */
func SearchKFromSlice(arr []int32, K, M int32) int32 {
	bitmap := make(map[int32]int32, 32)

	for i := 0; i < 32; i++ {
		k := int32(1 << i)
		bitmap[k] = int32(i)
	}

	t := make([]int32, 32)
	for _, v := range arr {
		for v != 0 {
			rightOne := v & (-v)
			t[ bitmap[rightOne] ]++
			v ^= rightOne
		}
	}

	ans := int32(0)

	for i, v := range t {
		if v % M != 0 {
			if v % M == K {
				ans |= 1 << i
			}else {
				return math.MinInt32
			}
		}
	}

	if ans == 0 {
		count := 0
		for _, v := range arr {
			if v == 0 {
				count++
			}
		}
		if int32(count) != K {
			return math.MinInt32
		}
	}

	return ans
}