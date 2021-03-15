package test

import (
	"fmt"
	"github.com/coredumptoday/practice/heap"
	"github.com/coredumptoday/practice/utils"
	"testing"
)

func TestCusHeap(t *testing.T) {
	value := 1000
	limit := 100
	testTimes := 1000000
	for i := 0; i < testTimes; i++ {
		curLimit := int(utils.GetRandNum()*float32(limit) + 1)
		my := heap.NewCusMaxHeap(curLimit)
		test := heap.NewHeapSlice(curLimit)
		curOpTimes := int(utils.GetRandNum() * float32(limit))
		for j := 0; j < curOpTimes; j++ {
			if my.IsEmpty() != test.IsEmpty() {
				fmt.Println("is empty Oops!")
			}
			if my.IsFull() != test.IsFull() {
				fmt.Println("is full Oops!")
			}
			if my.IsEmpty() {
				curValue := int(utils.GetRandNum() * float32(value))
				my.Push(curValue)
				test.Push(curValue)
			} else if my.IsFull() {
				pop1, ok1 := my.Pop()
				pop2, ok2 := test.Pop()

				if (ok1 == nil && ok2 != nil) || (ok1 != nil && ok2 == nil) {
					fmt.Println("is pop Oops1!", ok1, ok2)
					t.Fail()
					return
				} else if (ok1 == nil && ok2 == nil) && pop1 != pop2 {
					fmt.Println("is pop Oops1!", pop1, pop2)
					t.Fail()
					return
				}
			} else {
				if utils.GetRandNum() < 0.5 {
					curValue := int(utils.GetRandNum() * float32(value))
					my.Push(curValue)
					test.Push(curValue)
				} else {
					pop1, ok1 := my.Pop()
					pop2, ok2 := test.Pop()

					if (ok1 == nil && ok2 != nil) || (ok1 != nil && ok2 == nil) {
						fmt.Println("is pop Oops2!", ok1, ok2)
						t.Fail()
						return
					} else if (ok1 == nil && ok2 == nil) && pop1 != pop2 {
						fmt.Println("is pop Oops2!", pop1, pop2)
						t.Fail()
						return
					}
				}
			}
		}
	}
	fmt.Println("finish!")
}
