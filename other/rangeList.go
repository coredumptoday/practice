package other

import (
	"github.com/coredumptoday/practice/linear"
)

func RangeList(head *linear.Node, v int32) *linear.Node {
	if head == nil || head.Next == nil {
		return head
	}

	var lessH *linear.Node
	var lessT *linear.Node
	var equalH *linear.Node
	var equalT *linear.Node
	var moreH *linear.Node
	var moreT *linear.Node
	var next *linear.Node

	for head != nil {
		next = head.Next
		head.Next = nil
		if v == head.Value {
			if equalH == nil {
				equalH = head
				equalT = head
			} else {
				equalT.Next = head
				equalT = equalT.Next
			}
		} else if head.Value > v {
			if moreH == nil {
				moreH = head
				moreT = head
			} else {
				moreT.Next = head
				moreT = moreT.Next
			}
		} else {
			if lessH == nil {
				lessH = head
				lessT = head
			} else {
				lessT.Next = head
				lessT = lessT.Next
			}
		}

		head = next
	}

	if lessT != nil {
		lessT.Next = equalH
		if equalT == nil {
			equalT = lessT
		}
	}

	if equalT != nil {
		equalT.Next = moreH
	}

	if lessH != nil {
		return lessH
	} else if equalH != nil {
		return equalH
	}

	return moreH
}
