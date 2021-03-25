package test

import (
	"fmt"
	"testing"
)

/**
Maximum Subarray
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

题目要求
给定一个数组，求这个数组的连续子数组中的元素之和的最大值。

*/
func TestSubArrayTotalMax(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	max := arr[0]
	tmp := arr[0]

	for i := 1; i < len(arr); i++ {
		if tmp > 0 {
			tmp = tmp + arr[i]
		} else {
			tmp = arr[i]
		}
		if tmp > max {
			max = tmp
		}
	}

	fmt.Println()
	fmt.Println(max)
}
