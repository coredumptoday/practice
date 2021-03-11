package test

import (
	"fmt"
	"github.com/coredumptoday/practice/list"
	"github.com/coredumptoday/practice/utils"
	"testing"
)

func TestRevLinkedList(t *testing.T) {
	length := int32(50)
	value := int32(100)
	testTime := 100000
	fmt.Println("test begin!")
	for i := 0; i < testTime; i++ {
		linked := utils.GenerateRandomLinkedList(length, value)
		s := utils.GetLinkedListSlice(linked)
		linked = list.RevLinkedList(linked)

		if !utils.CheckSliceAndLinkedList(s, linked) {
			fmt.Print("LinkedList: ")
			utils.PrintLinkedList(linked)
			fmt.Print("Slice: ")
			utils.PrintSliceRev(s, false)
			t.Fail()
			break
		}
	}
	fmt.Println("test done!!!!")
}

func TestRevDoubleList(t *testing.T) {
	length := int32(50)
	value := int32(100)
	testTime := 100000
	fmt.Println("double test begin!")
	for i := 0; i < testTime; i++ {
		doubleList := utils.GenerateRandomDoubleList(length, value)
		s := utils.GetDoubleListSlice(doubleList)
		revList := list.RevDoubleList(doubleList)

		if !utils.CheckSliceAndDoubleList(s, revList) {
			fmt.Print("DoubleList: ")
			utils.PrintDoubleList(revList)
			fmt.Print("Slice: ")
			utils.PrintSliceRev(s, true)
			t.Fail()
			break
		}
	}
	fmt.Println("double test done!!!!")
}
