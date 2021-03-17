package tree

import (
	"container/list"
	"fmt"
)

type BinTreeNode struct {
	Value     int
	LeftNode  *BinTreeNode
	RightNode *BinTreeNode
}

func PrintBinTreePre(root *BinTreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Value)
	PrintBinTreePre(root.LeftNode)
	PrintBinTreePre(root.RightNode)
}

func PrintBinTreeMid(root *BinTreeNode) {
	if root == nil {
		return
	}
	PrintBinTreeMid(root.LeftNode)
	fmt.Print(root.Value)
	PrintBinTreeMid(root.RightNode)
}

func PrintBinTreePost(root *BinTreeNode) {
	if root == nil {
		return
	}
	PrintBinTreePost(root.LeftNode)
	PrintBinTreePost(root.RightNode)
	fmt.Print(root.Value)
}

func PrintBinTreePreNonRecursion(root *BinTreeNode) {
	if root == nil {
		return
	}
	stack := list.New()
	stack.PushFront(root)

	for stack.Len() > 0 {
		e := stack.Front()
		stack.Remove(e)

		cur := e.Value.(*BinTreeNode)
		fmt.Print(cur.Value)

		if cur.RightNode != nil {
			stack.PushFront(cur.RightNode)
		}
		if cur.LeftNode != nil {
			stack.PushFront(cur.LeftNode)
		}
	}
}

func PrintBinTreeMidNonRecursion(cur *BinTreeNode) {
	if cur == nil {
		return
	}

	stack := list.New()

	for stack.Len() > 0 || cur != nil {
		if cur != nil {
			stack.PushFront(cur)
			cur = cur.LeftNode
		} else {
			e := stack.Front()
			stack.Remove(e)
			cur = e.Value.(*BinTreeNode)
			fmt.Print(cur.Value)
			cur = cur.RightNode
		}
	}
}

func PrintBinTreePostNonRecursion(cur *BinTreeNode) {
	if cur == nil {
		return
	}

	stack := list.New()

	for stack.Len() > 0 || cur != nil {
		if cur != nil {
			stack.PushFront(cur)
			cur = cur.LeftNode
		} else {
			e := stack.Front()
			stack.Remove(e)
			cur = e.Value.(*BinTreeNode)
			fmt.Print(cur.Value)
			cur = cur.RightNode
		}
	}
}

func PrintBinTreeForDeep(cur *BinTreeNode) {
	if cur == nil {
		return
	}

	l := list.New()
	l.PushFront(cur)

	for l.Len() > 0 {
		e := l.Back()
		l.Remove(e)

		cur = e.Value.(*BinTreeNode)
		fmt.Print(cur.Value)

		if cur.LeftNode != nil {
			l.PushFront(cur.LeftNode)
		}

		if cur.RightNode != nil {
			l.PushFront(cur.RightNode)
		}
	}
}

func GetBinTreeMaxWidth(cur *BinTreeNode) int {
	if cur == nil {
		return 0
	}

	var curEnd, nextEnd *BinTreeNode
	l := list.New()
	l.PushFront(cur)
	curEnd = cur
	maxWidth := 0
	curWidth := 0

	for l.Len() > 0 {
		e := l.Back()
		l.Remove(e)
		cur = e.Value.(*BinTreeNode)
		curWidth++
		if cur.LeftNode != nil {
			l.PushFront(cur.LeftNode)
			nextEnd = cur.LeftNode
		}
		if cur.RightNode != nil {
			l.PushFront(cur.RightNode)
			nextEnd = cur.RightNode
		}

		if curEnd == cur {
			if curWidth > maxWidth {
				maxWidth = curWidth
			}
			curWidth = 0
			curEnd = nextEnd
			nextEnd = nil
		}
	}

	return maxWidth
}
