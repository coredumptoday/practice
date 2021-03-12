package utils

import (
	"fmt"
	"github.com/coredumptoday/practice/list"
)

func PrintLinkedList(head *list.Node) {
	for head != nil {
		fmt.Print(head.Value, " -> ")
		head = head.Next
	}
	fmt.Println("null")
}

func PrintDoubleList(head *list.DoubleNode) {
	for head != nil {
		if head.Prev == nil {
			fmt.Print("null <-")
		} else {
			fmt.Print("<-")
		}
		fmt.Print(head.Value)
		if head.Next == nil {
			fmt.Print("-> null")
		} else {
			fmt.Print("->")
		}
		head = head.Next
	}
	fmt.Println()
}

func GenerateRandomLinkedList(length, value int32) *list.Node {
	size := int32(GetRandNum() * float32(length+1))
	if size == 0 {
		return nil
	}

	head := list.Node{
		Value: int32(GetRandNum() * float32(value+1)),
		Next:  nil,
	}
	prev := &head
	size--

	for size != 0 {
		cur := list.Node{
			Value: int32(GetRandNum() * float32(value+1)),
			Next:  nil,
		}
		(*prev).Next = &cur
		prev = &cur
		size--
	}

	return &head
}

func GenerateRandomDoubleList(length, value int32) *list.DoubleNode {
	size := int32(GetRandNum() * float32(length+1))
	if size == 0 {
		return nil
	}

	head := list.DoubleNode{
		Value: int32(GetRandNum() * float32(value+1)),
		Next:  nil,
		Prev:  nil,
	}
	prev := &head
	size--

	for size != 0 {
		cur := list.DoubleNode{
			Value: int32(GetRandNum() * float32(value+1)),
			Next:  nil,
			Prev:  prev,
		}
		(*prev).Next = &cur
		prev = &cur
		size--
	}

	return &head
}

func GetLinkedListSlice(head *list.Node) []int32 {
	l := make([]int32, 0)
	for head != nil {
		l = append(l, head.Value)
		head = head.Next
	}
	return l
}

func GetDoubleListSlice(head *list.DoubleNode) []int32 {
	l := make([]int32, 0)
	for head != nil {
		l = append(l, head.Value)
		head = head.Next
	}
	return l
}

func PrintSliceRev(a []int32, isDouble bool) {
	t := ""
	if isDouble {
		t = " <-> "
	} else {
		t = " -> "
	}

	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(a[i], t)
	}
	fmt.Println("null")
}

func CheckSliceAndLinkedList(a []int32, head *list.Node) bool {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != head.Value {
			return false
		}
		head = head.Next
	}
	return true
}

func CheckSliceAndDoubleList(a []int32, head *list.DoubleNode) bool {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != head.Value {
			return false
		}
		head = head.Next
	}
	return true
}

func GetSliceWithNoValue(head *list.Node, v int32) []int32 {
	s := make([]int32, 0)
	for head != nil {
		if head.Value != v {
			s = append(s, head.Value)
		}
		head = head.Next
	}
	return s
}

func CheckRemoveValue(a []int32, head *list.Node) bool {
	for _, v := range a {
		if v != head.Value {
			return false
		}
		head = head.Next
	}
	return true
}