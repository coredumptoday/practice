package utils

import (
	"github.com/coredumptoday/practice/tree"
)

func generateBST(level, maxLevel, maxValue int) *tree.BinTreeNode {
	if level > maxValue || GetRandNum() < 0.5 {
		return nil
	}
	head := &tree.BinTreeNode{
		Value:     int(GetRandNum() * float32(maxValue)),
		LeftNode:  generateBST(level+1, maxLevel, maxValue),
		RightNode: generateBST(level+1, maxLevel, maxValue),
	}
	return head
}
func GenerateRandomBST(maxLevel, maxValue int) *tree.BinTreeNode {
	return generateBST(1, maxLevel, maxValue)
}

func in(cur *tree.BinTreeNode, arr []int) []int {
	if cur == nil {
		return arr
	}
	arr = in(cur.LeftNode, arr)
	arr = append(arr, cur.Value)
	arr = in(cur.RightNode, arr)
	return arr
}

func IsBST(cur *tree.BinTreeNode) bool {
	arr := make([]int, 0)
	arr = in(cur, arr)
	if len(arr) < 2 {
		return true
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}
