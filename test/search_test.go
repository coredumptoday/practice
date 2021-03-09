package test

import (
	"fmt"
	"github.com/coredumptoday/practice/search"
	"github.com/coredumptoday/practice/utils"
	"sort"
	"testing"
)

func TestExist(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		arr := utils.GenerateRandomSlice(maxSize, maxValue)
		sort.Ints(arr)
		value := utils.GetValue(maxValue)
		res := search.BinarySortExist(arr, value)
		if utils.IntInSlice(arr, value) != res {
			succeed = false
			fmt.Println("arr: ", arr)
			fmt.Println("value: ", value)
			t.Fail()
			break
		}

		if (i + 1) % 20000 == 0 {
			fmt.Println("BinarySortExist succeed count: ", i + 1)
			fmt.Println("res: ", res)
			fmt.Println("arr: ", arr)
			fmt.Println("value: ", value)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	}else {
		fmt.Println("Fail!")
	}
}