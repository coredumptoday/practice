package test

import (
	"fmt"
	"github.com/coredumptoday/practice/usort"
	"github.com/coredumptoday/practice/utils"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T)  {
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

		if (i + 1) % 20000 == 0 {
			fmt.Println("SelectionSort succeed count: ", i + 1)
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	}else {
		fmt.Println("Fail!")
	}
}

func TestBubbleSort(t *testing.T)  {
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

		if (i + 1) % 20000 == 0 {
			fmt.Println("BubbleSort succeed count: ", i + 1)
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	}else {
		fmt.Println("Fail!")
	}
}

func TestInsertionSort(t *testing.T)  {
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

		if (i + 1) % 20000 == 0 {
			fmt.Println("InsertionSort succeed count: ", i + 1)
			fmt.Println("a1: ", a1)
			fmt.Println("a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	}else {
		fmt.Println("Fail!")
	}
}