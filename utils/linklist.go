package utils

import (
	"fmt"
	"github.com/coredumptoday/practice/linear"
)

func PrintLinkedList(head *linear.Node) {
	for head != nil {
		fmt.Print(head.Value)
		head = head.Next
		if head != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println("")
}

func PrintRandLinkedList(head *linear.NodeJmp) {
	cur := head
	fmt.Println("next: ")
	for cur != nil {
		fmt.Print(cur.Value)
		cur = cur.Next
		if cur != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println("")
	fmt.Println("jump: ")
	for head != nil {
		fmt.Print(head.Value)
		head = head.Jmp
		if head != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println("")
}

func PrintLinkedListPos(head *linear.Node, n *linear.Node) {
	for head != nil {
		if head == n {
			fmt.Print("^", "    ")
		} else {
			fmt.Print(" ", "    ")
		}

		head = head.Next
	}
	fmt.Println("")
}

func PrintDoubleList(head *linear.DoubleNode) {
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

func GenerateRandomLinkedList(length, value int32) *linear.Node {
	size := int32(GetRandNum() * float32(length+1))
	if size == 0 {
		return nil
	}

	head := linear.Node{
		Value: int32(GetRandNum() * float32(value+1)),
		Next:  nil,
	}
	prev := &head
	size--

	for size != 0 {
		cur := linear.Node{
			Value: int32(GetRandNum() * float32(value+1)),
			Next:  nil,
		}
		(*prev).Next = &cur
		prev = &cur
		size--
	}

	return &head
}

func GenerateRandomDoubleList(length, value int32) *linear.DoubleNode {
	size := int32(GetRandNum() * float32(length+1))
	if size == 0 {
		return nil
	}

	head := linear.DoubleNode{
		Value: int32(GetRandNum() * float32(value+1)),
		Next:  nil,
		Prev:  nil,
	}
	prev := &head
	size--

	for size != 0 {
		cur := linear.DoubleNode{
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

func GetLinkedListSlice(head *linear.Node) []int32 {
	l := make([]int32, 0)
	for head != nil {
		l = append(l, head.Value)
		head = head.Next
	}
	return l
}

func GetDoubleListSlice(head *linear.DoubleNode) []int32 {
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

func CheckSliceAndLinkedList(a []int32, head *linear.Node) bool {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != head.Value {
			return false
		}
		head = head.Next
	}
	return true
}

func CheckSliceAndDoubleList(a []int32, head *linear.DoubleNode) bool {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != head.Value {
			return false
		}
		head = head.Next
	}
	return true
}

func GetSliceWithNoValue(head *linear.Node, v int32) []int32 {
	s := make([]int32, 0)
	for head != nil {
		if head.Value != v {
			s = append(s, head.Value)
		}
		head = head.Next
	}
	return s
}

func CheckRemoveValue(a []int32, head *linear.Node) bool {
	for _, v := range a {
		if v != head.Value {
			return false
		}
		head = head.Next
	}
	return true
}
