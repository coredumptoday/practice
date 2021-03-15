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
