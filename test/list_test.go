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
	fmt.Println("test begin!")
	for i := 0; i < testTime; i++ {
		v := int32(utils.GetRandNum() * float32(value+1))
		linked := utils.GenerateRandomLinkedList(length, value)
		s := utils.GetSliceWithNoValue(linked, v)
		linked = list.RemoveValue(linked, v)

		if !utils.CheckRemoveValue(s, linked) {
			fmt.Println("RemoveValue: ", v)
			fmt.Print("LinkedList: ")
			utils.PrintLinkedList(linked)
			fmt.Println("Slice: ", s)
			t.Fail()
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("Remove Check succeed count: ", i+1)
			fmt.Println("RemoveValue: ", v)
			fmt.Print("LinkedList: ")
			utils.PrintLinkedList(linked)
			fmt.Println("Slice: ", s)
		}
	}
	fmt.Println("test done!!!")
}

func TestQueue(t *testing.T) {
	oneTestDataNum := 100
	value := 10000
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		myQueue := list.NewQueue()
		stdQueue := clist.New()

		for j := 0; j < oneTestDataNum; j++ {
			num := int32(utils.GetRandNum() * float32(value))

			if stdQueue.Len() == 0 {
				myQueue.Push(num)
				stdQueue.PushFront(num)
			} else {
				if utils.GetRandNum() < 0.5 {
					myQueue.Push(num)
					stdQueue.PushFront(num)
				} else {
					myValue := myQueue.Pop()
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

		if myQueue.Len() != int32(stdQueue.Len()) {
			fmt.Println("length not equal")
			t.Fail()
			return
		}
	}
}

func TestStack(t *testing.T) {
	oneTestDataNum := 100
	value := 10000
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		myStack := list.NewStack()
		stdStack := clist.New()

		for j := 0; j < oneTestDataNum; j++ {
			num := int32(utils.GetRandNum() * float32(value))

			if stdStack.Len() == 0 {
				myStack.Push(num)
				stdStack.PushFront(num)
			} else {
				if utils.GetRandNum() < 0.5 {
					myStack.Push(num)
					stdStack.PushFront(num)
				} else {
					myValue := myStack.Pop()
					stdValue := stdStack.Front()
					stdStack.Remove(stdValue)

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

		if myStack.Len() != int32(stdStack.Len()) {
			fmt.Println("length not equal")
			t.Fail()
			return
		}
	}
}

func TestRingArray(t *testing.T) {
	oneTestDataNum := 100
	value := 10000
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		cap := int(utils.GetRandNum() * float32(value))
		ringArr := list.NewRingArr(cap)
		stdQueue := clist.New()

		for j := 0; j < oneTestDataNum; j++ {
			num := int(utils.GetRandNum() * float32(value))

			if ringArr.Len() == 0 {
				if ringArr.Push(num) {
					stdQueue.PushFront(num)
				}
			} else {
				if utils.GetRandNum() < 0.5 {
					if ringArr.Push(num) {
						stdQueue.PushFront(num)
					}
				} else {
					myValue, ok := ringArr.Pop()
					stdValue := stdQueue.Back()
					stdQueue.Remove(stdValue)

					if (ok == false && stdValue != nil) || (ok != false && stdValue == nil) {
						fmt.Println("oops!")
						fmt.Println(myValue, stdValue)
						t.Fail()
						return
					} else if (ok && stdValue != nil) && myValue != stdValue.Value {
						fmt.Println("oops!")
						fmt.Println(myValue, stdValue)
						t.Fail()
						return
					}
				}
			}
		}

		if ringArr.Len() != stdQueue.Len() {
			fmt.Println("length not equal")
			fmt.Println(ringArr.Len(), stdQueue.Len(), ringArr, stdQueue)
			t.Fail()
			return
		}
	}
}
