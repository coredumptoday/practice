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

	h1 := head
	h2 := head

	for h2.Next != nil && h2.Next.Next != nil {
		h1 = h1.Next
		h2 = h2.Next.Next
	}

	h2 = h1.Next  //h2指向后半段起始
	h1.Next = nil //前半段尾部指nil

	h3 := h1
	var next *linear.Node

	for h2 != nil {
		next = h2.Next
		h2.Next = h3
		h3 = h2
		h2 = next
	}

	h2 = head
	revh := h3

	res := true
	for h2 != nil && revh != nil {
		if h2.Value != revh.Value {
			res = false
			break
		}
		h2 = h2.Next
		revh = revh.Next
	}

	revh = h3
	h3 = nil
	next = nil

	for revh != nil {
		next = revh.Next
		revh.Next = h3
		h3 = revh
		revh = next
	}

	//h1.Next = h3
	return res
}
