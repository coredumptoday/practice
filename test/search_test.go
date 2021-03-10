package test

import (
	"fmt"
	"github.com/coredumptoday/practice/search"
	"github.com/coredumptoday/practice/utils"
	"github.com/coredumptoday/practice/xor"
	"math"
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

func TestSearchLeft(t *testing.T) {
	testTime := 500000
	maxSize := 40
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		arr := utils.GenerateRandomSlice(maxSize, maxValue)
		sort.Ints(arr)
		value := utils.GetValue(maxValue)
		idx := search.BinarySortSearchLeft(arr, value)

		idx2 := -1
		for i, v := range arr {
			if v >= value {
				idx2 = i
				break
			}
		}

		if idx2 != idx {
			succeed = false
			fmt.Println("idx: ", idx)
			fmt.Println("arr: ", arr)
			fmt.Println("value: ", value)
			t.Fail()
			break
		}

		if (i + 1) % 20000 == 0 {
			fmt.Println("BinarySortSearchLeft succeed count: ", i + 1)
			fmt.Println("idx: ", idx)
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

func TestSearchRight(t *testing.T) {
	testTime := 500000
	maxSize := 40
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		arr := utils.GenerateRandomSlice(maxSize, maxValue)
		sort.Ints(arr)
		value := utils.GetValue(maxValue)
		idx := search.BinarySortSearchRight(arr, value)

		idx2 := -1
		for i, v := range arr {
			if v > value {
				break
			}else {
				idx2 = i
			}
		}

		if idx2 != idx {
			succeed = false
			fmt.Println("idx: ", idx)
			fmt.Println("arr: ", arr)
			fmt.Println("value: ", value)
			t.Fail()
			break
		}

		if (i + 1) % 20000 == 0 {
			fmt.Println("BinarySortSearchLeft succeed count: ", i + 1)
			fmt.Println("idx: ", idx)
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

func TestOnlyOneOddNum(t *testing.T) {
	arr := []int{3, 3, 2, 3, 1, 1, 1, 3, 1, 1, 1}
	fmt.Println(xor.SearchOnlyOneShowOddNumFromSlice(arr))
}

func Test2OddNum(t *testing.T) {
	arr := []int{ 4, 3, 4, 2, 2, 2, 4, 1, 1, 1, 3, 3, 1, 1, 1, 4, 2, 2 }
	fmt.Println(xor.SearchTwoShowOddNumFromSlice(arr))
}

func TestSearchKM(t *testing.T) {
	kinds := 5
	xrange := 30
	testTime := 100000
	max := 9
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		a := int( utils.GetRandNum() * float32(max) + 1 ) // a 1 ~ 9
		b := int( utils.GetRandNum() * float32(max) + 1 ) // b 1 ~ 9
		k := int( math.Min(float64(a), float64(b)) )
		m := int( math.Max(float64(a), float64(b)) )
		// k < m
		if k == m {
			m++
		}

		arr, kTimeNum := utils.MKRandomSlice(kinds, xrange, k, m)
		ans1 := utils.KmTest(arr, int32(k))
		ans2 := xor.SearchKFromSlice(arr, int32(k), int32(m))
		if ans1 != ans2 {
			fmt.Printf("arr: %v, k: %v, kNum: %v\n", arr, k, kTimeNum)
			fmt.Println(ans1)
			fmt.Println(ans2)
			fmt.Println("出错了！")
			t.Fail()
			break
		}
	}
	fmt.Println("测试结束")
}