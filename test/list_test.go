package test

import (
	clist "container/list"
	"fmt"
	"github.com/coredumptoday/practice/linear"
	"github.com/coredumptoday/practice/utils"
	"testing"
)

//单向链表翻转
func TestRevLinkedList(t *testing.T) {
	length := int32(50)
	value := int32(100)
	testTime := 100000
	fmt.Println("test begin!")
	for i := 0; i < testTime; i++ {
		linked := utils.GenerateRandomLinkedList(length, value)
		s := utils.GetLinkedListSlice(linked)
		linked = linear.RevLinkedList(linked)

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

//双向链表翻转
func TestRevDoubleList(t *testing.T) {
	length := int32(50)
	value := int32(100)
	testTime := 100000
	fmt.Println("double test begin!")
	for i := 0; i < testTime; i++ {
		doubleList := utils.GenerateRandomDoubleList(length, value)
		s := utils.GetDoubleListSlice(doubleList)
		revList := linear.RevDoubleList(doubleList)

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

//从链表中删除值为value的元素
func TestRemoveValue(t *testing.T) {
	length := int32(50)
	value := int32(100)
	testTime := 100000
	fmt.Println("test begin!")
	for i := 0; i < testTime; i++ {
		v := int32(utils.GetRandNum() * float32(value+1))
		linked := utils.GenerateRandomLinkedList(length, value)
		s := utils.GetSliceWithNoValue(linked, v)
		linked = linear.RemoveValue(linked, v)

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

//双向链表手写队列
func TestQueue(t *testing.T) {
	oneTestDataNum := 100
	value := 10000
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		myQueue := linear.NewQueue()
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

//双向链表手写栈
func TestStack(t *testing.T) {
	oneTestDataNum := 100
	value := 10000
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		myStack := linear.NewStack()
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

//数组环形队列
func TestRingArray(t *testing.T) {
	oneTestDataNum := 100
	value := 10000
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		cap := int(utils.GetRandNum() * float32(value))
		ringArr := linear.NewRingArr(cap)
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

//实现栈结构，并且O(1)返回最小值
func TestMinStack(t *testing.T) {
	mStack := linear.NewMinStack()

	fmt.Println("push", 3)
	mStack.Push(3)
	fmt.Println("getMin", mStack.GetMin())
	mStack.Push(4)
	fmt.Println("push", 4)
	fmt.Println("getMin", mStack.GetMin())
	mStack.Push(1)
	fmt.Println("push", 1)
	fmt.Println("getMin", mStack.GetMin())
	fmt.Println("pop", mStack.Pop())
	fmt.Println("getMin", mStack.GetMin())
}

//用栈结构实现队列
func TestStackQueue(t *testing.T) {
	sQueue := linear.NewStackQueue()
	sQueue.Add(1)
	sQueue.Add(2)
	sQueue.Add(3)

	a, ok := sQueue.Peek()
	fmt.Println("peek", a, ok)

	a, ok = sQueue.Poll()
	fmt.Println("poll", a, ok)

	a, ok = sQueue.Peek()
	fmt.Println("peek", a, ok)

	a, ok = sQueue.Poll()
	fmt.Println("poll", a, ok)

	a, ok = sQueue.Peek()
	fmt.Println("peek", a, ok)

	a, ok = sQueue.Poll()
	fmt.Println("poll", a, ok)

	a, ok = sQueue.Peek()
	fmt.Println("peek", a, ok)
}

//用队列结构实现栈
func TestQueueStack(t *testing.T) {
	fmt.Println("test begin")
	myStack := linear.NewQueueStack()
	stdStack := clist.New()
	testTime := 1000000
	max := 1000000
	for i := 0; i < testTime; i++ {
		if myStack.IsEmpty() {
			if !(stdStack.Len() == 0) {
				fmt.Println("empty Oops", stdStack.Len(), myStack.Len())
				t.Fail()
				return
			}
			num := int(utils.GetRandNum() * float32(max))
			myStack.Push(num)
			stdStack.PushFront(num)
		} else {
			if utils.GetRandNum() < 0.25 {
				num := int(utils.GetRandNum() * float32(max))
				myStack.Push(num)
				stdStack.PushFront(num)
			} else if utils.GetRandNum() < 0.5 {
				v, ok := myStack.Peek()
				stdV := stdStack.Front()

				if (ok && stdV == nil) || (!ok && stdV != nil) {
					fmt.Println("peek1 Oops")
					t.Fail()
					return
				} else if (ok && stdV != nil) && stdV.Value != v {
					fmt.Println("peek2 Oops", ok, stdV, stdV.Value, v)
					t.Fail()
					return
				}
			} else if utils.GetRandNum() < 0.75 {
				v, ok := myStack.Poll()
				stdV := stdStack.Front()
				if stdV != nil {
					stdStack.Remove(stdV)
				}

				if (ok && stdV == nil) || (!ok && stdV != nil) {
					fmt.Println("poll1 Oops")
					t.Fail()
					return
				} else if ok && stdV != nil && stdV.Value != v {
					fmt.Println("poll2 Oops", ok, stdV, stdV.Value, v)
					t.Fail()
					return
				}
			} else {
				if myStack.IsEmpty() != (stdStack.Len() == 0) {
					fmt.Println("Oops", stdStack.Len(), myStack.Len())
				}
			}
		}
	}

	fmt.Println("test finish!")
}

//在m-n的范围上翻转单链表
func TestRevLinkListForRange(t *testing.T) {
	l := &linear.Node{}
	l.Next = &linear.Node{Value: 1, Next: nil}
	l.Next.Next = &linear.Node{Value: 2, Next: nil}
	l.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	l.Next.Next.Next.Next = &linear.Node{Value: 4, Next: nil}
	l.Next.Next.Next.Next.Next = &linear.Node{Value: 5, Next: nil}
	l.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 6, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 7, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 8, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 9, Next: nil}

	utils.PrintLinkedList(l)
	linear.RevLinkedListForRange(l, 3, 7)
	utils.PrintLinkedList(l)
}

//每K个元素翻转一次单链表
func TestRevLinkListForGroupK(t *testing.T) {
	l := &linear.Node{}
	l.Next = &linear.Node{Value: 1, Next: nil}
	l.Next.Next = &linear.Node{Value: 2, Next: nil}
	l.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	l.Next.Next.Next.Next = &linear.Node{Value: 4, Next: nil}
	l.Next.Next.Next.Next.Next = &linear.Node{Value: 5, Next: nil}
	l.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 6, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 7, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 8, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 9, Next: nil}

	utils.PrintLinkedList(l)
	a := linear.RevLinkedListForGroupK(l, 3)
	utils.PrintLinkedList(a)
}

//删除链表m-n的元素
func TestLinkListRemoveRange(t *testing.T) {
	l := &linear.Node{}
	l.Next = &linear.Node{Value: 1, Next: nil}
	l.Next.Next = &linear.Node{Value: 2, Next: nil}
	l.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	l.Next.Next.Next.Next = &linear.Node{Value: 4, Next: nil}
	l.Next.Next.Next.Next.Next = &linear.Node{Value: 5, Next: nil}
	l.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 6, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 7, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 8, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 9, Next: nil}

	utils.PrintLinkedList(l)
	linear.LinkListRemoveRange(l, 3, 5)
	utils.PrintLinkedList(l)
}

func TestRemoveDupNodeFromSortList(t *testing.T) {
	l := &linear.Node{}
	l.Next = &linear.Node{Value: 1, Next: nil}
	l.Next.Next = &linear.Node{Value: 2, Next: nil}
	l.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	l.Next.Next.Next.Next = &linear.Node{Value: 4, Next: nil}
	l.Next.Next.Next.Next.Next = &linear.Node{Value: 5, Next: nil}
	l.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 6, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 7, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 8, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 9, Next: nil}

	utils.PrintLinkedList(l)
	linear.RemoveDupNodeFromSortList(l)
	utils.PrintLinkedList(l)

	l = &linear.Node{}
	l.Next = &linear.Node{Value: 1, Next: nil}
	l.Next.Next = &linear.Node{Value: 1, Next: nil}
	l.Next.Next.Next = &linear.Node{Value: 1, Next: nil}
	l.Next.Next.Next.Next = &linear.Node{Value: 2, Next: nil}
	l.Next.Next.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	l.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 8, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 9, Next: nil}

	utils.PrintLinkedList(l)
	linear.RemoveDupNodeFromSortList(l)
	utils.PrintLinkedList(l)
}
