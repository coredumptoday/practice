package test

import (
	"fmt"
	"github.com/coredumptoday/practice/heap"
	"github.com/coredumptoday/practice/usort"
	"github.com/coredumptoday/practice/utils"
	"sort"
	"testing"
)

//选择排序
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

//冒泡排序
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

//插入排序
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
			h.Push(v)
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
