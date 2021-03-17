package other

import (
	"github.com/coredumptoday/practice/linear"
)

func CopyLinkListWithRandPtr(head *linear.NodeJmp) *linear.NodeJmp {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head

	for cur != nil {
		nNode := &linear.NodeJmp{
			Value: cur.Value,
			Next:  cur.Next,
		}
		cur.Next = nNode
		cur = nNode.Next
	}

	cur = head
	for cur != nil {
		if cur.Jmp == nil {
			cur.Next.Jmp = nil
		} else {
			cur.Next.Jmp = cur.Jmp.Next
		}

		cur = cur.Next.Next
	}

	var n = head.Next

	cur = head
	for cur != nil {
		next := cur.Next.Next
		cp := cur.Next
		cur.Next = next

		if next != nil {
			cp.Next = next.Next
		} else {
			cp.Next = nil
		}

		cur = cur.Next
	}

	return n
}
