package test

import (
	"fmt"
	"github.com/coredumptoday/practice/other"
	"github.com/coredumptoday/practice/utils"
	"testing"
)

func TestLessSum(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		a1 := utils.GenerateRandomSlice(maxSize, maxValue)
		a2 := make([]int, len(a1))
		copy(a2, a1)

		s1 := other.LessSum(a1)
		s2 := other.Comparator(a2)

		if s1 != s2 {
			succeed = false
			fmt.Println("sum1: ", s1, "a1: ", a1)
			fmt.Println("sum2: ", s2, "a2: ", a2)
			t.Fail()
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("lessSum succeed count: ", i+1)
			fmt.Println("sum1: ", s1, "a1: ", a1)
			fmt.Println("sum2: ", s2, "a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fail!")
	}
}
