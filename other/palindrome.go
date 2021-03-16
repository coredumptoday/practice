package other

import (
	"container/list"
	"github.com/coredumptoday/practice/linear"
)

func IsPalindromeWithNSpace(head *linear.Node) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	stack := list.New()
	cur := head
	for cur != nil {
		stack.PushFront(cur.Value)
		cur = cur.Next
	}

	for head != nil {
		e := stack.Front()
		stack.Remove(e)

		v := e.Value.(int32)
		if head.Value != v {
			return false
		}
		head = head.Next
	}

	return true
}

func IsPalindromeWithHalfNSpace(head *linear.Node) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	stack := list.New()

	slow := head.Next
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	for slow != nil {
		stack.PushFront(slow.Value)
		slow = slow.Next
	}

	for stack.Len() > 0 {
		e := stack.Front()
		stack.Remove(e)

		if e.Value != head.Value {
			return false
		}
		head = head.Next
	}
	return true
}

func IsPalindromeWith1Space(head *linear.Node) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	n1 := head
	n2 := head

	for n2.Next != nil && n2.Next.Next != nil {
		n1 = n1.Next
		n2 = n2.Next.Next
	}

	n2 = n1.Next
	n1.Next = nil

	var n3 *linear.Node

	for n2 != nil {
		n3 = n2.Next
		n2.Next = n1
		n1 = n2
		n2 = n3
	}

	n3 = n1
	n2 = head
	res := true

	for n1 != nil && n2 != nil {
		if n1.Value != n2.Value {
			res = false
			break
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	n1 = n3.Next
	n3.Next = nil
	for n1 != nil { // recover list
		n2 = n1.Next
		n1.Next = n3
		n3 = n1
		n1 = n2
	}
	return res
}
