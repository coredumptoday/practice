package test

import (
	"fmt"
	"github.com/coredumptoday/practice/tree"
	"github.com/coredumptoday/practice/utils"
	"testing"
)

func TestPrefixTree(t *testing.T) {
	arrLen := 100
	strLen := 20
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		arr := utils.GenerateRandomStringArray(arrLen, strLen)
		preTree := tree.NewPrefixTree()
		right := tree.NewPrefixRight()
		for j := 0; j < len(arr); j++ {
			decide := utils.GetRandNum()
			if decide < 0.25 {
				preTree.Insert(arr[j])
				right.Insert(arr[j])
			} else if decide < 0.5 {
				preTree.Delete(arr[j])
				right.Delete(arr[j])
			} else if decide < 0.75 {
				ans1 := preTree.Search(arr[j])
				ans2 := right.Search(arr[j])
				if ans1 != ans2 {
					fmt.Println("Search Oops!", ans1, ans2)
					t.Fail()
					return
				}
			} else {
				ans1 := preTree.PrefixCount(arr[j])
				ans2 := right.PrefixCount(arr[j])
				if ans1 != ans2 {
					fmt.Println("PrefixCount Oops!", ans1, ans2, arr[j], arr, j)
					t.Fail()
					return
				}
			}
		}
	}
	fmt.Println("finish!")

}

func TestBinTreePrint(t *testing.T) {
	root := &tree.BinTreeNode{Value: 1}
	root.LeftNode = &tree.BinTreeNode{Value: 2}
	root.RightNode = &tree.BinTreeNode{Value: 3}
	root.LeftNode.LeftNode = &tree.BinTreeNode{Value: 4}
	root.LeftNode.RightNode = &tree.BinTreeNode{Value: 5}
	root.RightNode.LeftNode = &tree.BinTreeNode{Value: 6}
	root.RightNode.RightNode = &tree.BinTreeNode{Value: 7}

	tree.PrintBinTreePre(root)
	fmt.Println("========先序递归")
	tree.PrintBinTreePreNonRecursion(root)
	fmt.Println("========先序非递归")
	tree.PrintBinTreeMid(root)
	fmt.Println("========中序递归")
	tree.PrintBinTreeMidNonRecursion(root)
	fmt.Println("========中序非递归")
	tree.PrintBinTreePost(root)
	fmt.Println("========后续递归")
	tree.PrintBinTreeForDeep(root)
	fmt.Println("========层序")

}

func TestTreeMaxWidth(t *testing.T) {
	root := &tree.BinTreeNode{Value: 1}
	root.LeftNode = &tree.BinTreeNode{Value: 2}
	root.RightNode = &tree.BinTreeNode{Value: 3}
	root.LeftNode.LeftNode = &tree.BinTreeNode{Value: 4}
	root.LeftNode.RightNode = &tree.BinTreeNode{Value: 5}
	root.RightNode.LeftNode = &tree.BinTreeNode{Value: 6}
	root.RightNode.RightNode = &tree.BinTreeNode{Value: 7}
	root.LeftNode.RightNode.LeftNode = &tree.BinTreeNode{Value: 8}
	tree.PrintBinTreeForDeep(root)
	fmt.Println("========层序")
	fmt.Println(tree.GetBinTreeMaxWidth(root))
}

func TestTreeComplete(t *testing.T) {
	root := &tree.BinTreeNode{Value: 1}
	root.LeftNode = &tree.BinTreeNode{Value: 2}
	root.RightNode = &tree.BinTreeNode{Value: 3}
	root.LeftNode.LeftNode = &tree.BinTreeNode{Value: 4}
	root.LeftNode.RightNode = &tree.BinTreeNode{Value: 5}
	root.RightNode.LeftNode = &tree.BinTreeNode{Value: 6}
	root.RightNode.RightNode = &tree.BinTreeNode{Value: 7}
	root.LeftNode.RightNode.LeftNode = &tree.BinTreeNode{Value: 8}
	tree.PrintBinTreeForDeep(root)
	fmt.Println("========层序")
	fmt.Println(tree.IsCompleteBinaryTree(root))

	root = &tree.BinTreeNode{Value: 1}
	root.LeftNode = &tree.BinTreeNode{Value: 2}
	root.RightNode = &tree.BinTreeNode{Value: 3}
	root.LeftNode.LeftNode = &tree.BinTreeNode{Value: 4}
	root.LeftNode.RightNode = &tree.BinTreeNode{Value: 5}
	root.RightNode.LeftNode = &tree.BinTreeNode{Value: 6}
	root.RightNode.RightNode = &tree.BinTreeNode{Value: 7}

	tree.PrintBinTreeForDeep(root)
	fmt.Println("========层序")
	fmt.Println(tree.IsCompleteBinaryTree(root))
}
