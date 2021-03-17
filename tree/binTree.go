package tree

import "fmt"

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
