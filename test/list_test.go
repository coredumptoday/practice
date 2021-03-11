package test

import (
	clist "container/list"
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

func TestRemoveValue(t *testing.T) {
	length := int32(50)
	value := int32(100)
	testTime := 100000
	v := int32(utils.GetRandNum() * float32(value+1))
	fmt.Println("test begin!")
	for i := 0; i < testTime; i++ {
		linked := utils.GenerateRandomLinkedList(length, value)
		s := utils.GetSliceWithNoValue(linked, v)

		if !utils.CheckRemoveValue(s, linked) {
			fmt.Println("RemoveValue: ", v)
			fmt.Print("LinkedList: ")
			utils.PrintLinkedList(linked)
			fmt.Print("Slice: ")
			utils.PrintSliceRev(s, false)
			t.Fail()
			break
		}
	}
}

func TestQueue(t *testing.T) {
	oneTestDataNum := 100
	value := 10000
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		myQueue := list.DoubleEndsQueue{}
		stdQueue := clist.New()

		for j := 0; j < oneTestDataNum; j++ {
			num := int32(utils.GetRandNum() * float32(value))

			if stdQueue.Len() == 0 {
				myQueue.AddFromHead(num)
				stdQueue.PushFront(num)
			} else {
				if utils.GetRandNum() < 0.5 {
					myQueue.AddFromHead(num)
					stdQueue.PushFront(num)
				} else {
					myValue := myQueue.PopFromBottom()
					stdValue := stdQueue.Back()
					stdQueue.Remove(stdValue)

					if (myValue == nil && stdValue != nil) || (myValue != nil && stdValue == nil) {
						fmt.Println("oops!")
						fmt.Println(myValue, stdValue)
						t.Fail()
						return
					} else if (myValue != nil && stdValue != nil) && myValue.Value != stdValue.Value {
						fmt.Println("oops!")
						fmt.Println(myValue, stdValue)
						t.Fail()
						return
					}
				}
			}
		}

		if myQueue.GetLength() != int32(stdQueue.Len()) {
			fmt.Println("length not equal")
			t.Fail()
			return
		}
	}
}
