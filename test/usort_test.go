package test

import (
	"fmt"
	"github.com/coredumptoday/practice/heap"
	"github.com/coredumptoday/practice/usort"
	"github.com/coredumptoday/practice/utils"
	"sort"
	"testing"
)

/*
	选择排序
	1. 从 0 ~ N-1 位置选择一个最小的放在 0 位置
	2. 从 1 ~ N-1 位置选择一个最小的放在 1 位置
	3. 从 2 ~ N-1 位置选择一个最小的放在 2 位置
	...
	...
*/
func TestSelectionSort(t *testing.T) {
	testTime := 500000
	maxSize := 50
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		a1 := utils.GenerateRandomSlice(maxSize, maxValue)
		a2 := make([]int, len(a1))
		copy(a2, a1)

		usort.SelectionSort(a1)
		sort.Ints(a2)

		if !utils.IsSliceEqual(a1, a2) {
			succeed = false
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("SelectionSort succeed count: ", i+1)
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fail!")
	}
}

/*
	冒泡排序
	1. 从 0 ~ N-1 位置中选择一个最大的放在 N-1 位置
	2. 从 0 ~ N-2 位置中选择一个最大的放在 N-2 位置
	3. 从 0 ~ N-3 位置中选择一个最大的放在 N-3 位置
	...
	...
*/
func TestBubbleSort(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		a1 := utils.GenerateRandomSlice(maxSize, maxValue)
		a2 := make([]int, len(a1))
		copy(a2, a1)

		usort.BubbleSort(a1)
		sort.Ints(a2)

		if !utils.IsSliceEqual(a1, a2) {
			succeed = false
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("BubbleSort succeed count: ", i+1)
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fail!")
	}
}

/*
	插入排序
	1. 0 位置上有序
	2. 0 ~ 1 位置上有序
	3. 0 ~ 2 位置上有序
	4. 0 ~ 3 位置上有序
	....
	...
*/
func TestInsertionSort(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		a1 := utils.GenerateRandomSlice(maxSize, maxValue)
		a2 := make([]int, len(a1))
		copy(a2, a1)

		usort.InsertionSort(a1)
		sort.Ints(a2)

		if !utils.IsSliceEqual(a1, a2) {
			succeed = false
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("InsertionSort succeed count: ", i+1)
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fail!")
	}
}

//归并排序
func TestMergeSort(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		a1 := utils.GenerateRandomSlice(maxSize, maxValue)
		a2 := make([]int, len(a1))
		copy(a2, a1)

		usort.MergeSort(a1)
		sort.Ints(a2)

		if !utils.IsSliceEqual(a1, a2) {
			succeed = false
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("MergeSort succeed count: ", i+1)
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fail!")
	}
}

//快速排序
func TestQuickSort(t *testing.T) {
	testTime := 500000
	maxSize := 50
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		a1 := utils.GenerateRandomSlice(maxSize, maxValue)
		a2 := make([]int, len(a1))
		a3 := make([]int, len(a1))
		a4 := make([]int, len(a1))
		a5 := make([]int, len(a1))
		copy(a2, a1)
		copy(a3, a1)
		copy(a4, a1)
		copy(a5, a1)

		usort.QuickSortWith2Range(a1)
		usort.QuickSortWith3Range(a2)
		usort.QuickSortWithRandNum(a3)
		usort.QuickSortNonRecursion(a4)
		sort.Ints(a5)

		if !utils.IsSliceEqual(a1, a5) || !utils.IsSliceEqual(a2, a5) || !utils.IsSliceEqual(a3, a5) || !utils.IsSliceEqual(a4, a5) {
			succeed = false
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
			fmt.Println("a3: ", a3)
			fmt.Println("a4: ", a4)
			fmt.Println("a5: ", a5)
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("HeapSort succeed count: ", i+1)
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
			fmt.Println("a3: ", a3)
			fmt.Println("a4: ", a4)
			fmt.Println("a5: ", a5)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fail!")
	}
}

//堆排序
func TestHeapSort(t *testing.T) {
	testTime := 500000
	maxSize := 50
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		a1 := utils.GenerateRandomSlice(maxSize, maxValue)
		a2 := make([]int, len(a1))
		a3 := make([]int, len(a1))
		copy(a2, a1)
		copy(a3, a1)

		h := heap.NewCusMaxHeap(len(a3))
		for _, v := range a3 {
			_ = h.Push(v)
		}
		for i := len(a3) - 1; i >= 0; i-- {
			v, _ := h.Pop()
			a3[i] = v
		}

		usort.HeapSort(a1)
		sort.Ints(a2)

		if !utils.IsSliceEqual(a1, a2) || !utils.IsSliceEqual(a3, a2) {
			succeed = false
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
			fmt.Println("a3: ", a3)
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("HeapSort succeed count: ", i+1)
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
			fmt.Println("a3: ", a3)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fail!")
	}
}

//计数排序
func TestCountSort(t *testing.T) {
	testTime := 500000
	maxSize := 50
	maxValue := 150
	succeed := true

	for i := 0; i < testTime; i++ {
		a1 := utils.GeneratePosRandomSlice(maxSize, maxValue)
		a2 := make([]int, len(a1))
		copy(a2, a1)

		usort.CountSort(a1, maxValue)
		sort.Ints(a2)

		if !utils.IsSliceEqual(a1, a2) {
			succeed = false
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("HeapSort succeed count: ", i+1)
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fail!")
	}
}
