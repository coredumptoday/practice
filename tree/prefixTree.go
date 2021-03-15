package tree

import "strings"

var CharArr = []string{
	"a", "b", "c", "d", "e",
	"f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o",
	"p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y",
	"z",
}

func newNode() *prefixTreeNode {
	m := make(map[string]*prefixTreeNode, len(CharArr))
	for _, v := range CharArr {
		m[v] = nil
	}

	return &prefixTreeNode{
		passCount: 0,
		endCount:  0,
		path:      m,
	}
}

type prefixTreeNode struct {
	passCount int
	endCount  int
	path      map[string]*prefixTreeNode
}

func NewPrefixTree() *PrefixTree {
	return &PrefixTree{
		root: newNode(),
	}
}

type PrefixTree struct {
	root *prefixTreeNode
}

func (tree *PrefixTree) Insert(str string) {
	tree.root.passCount++
	if str == "" {
		tree.root.endCount++
	}

	cArr := strings.Split(str, "")
	curNode := tree.root

	for i, v := range cArr {
		if curNode.path[v] == nil {
			curNode.path[v] = newNode()
		}
		curNode.path[v].passCount++
		if i == len(cArr)-1 {
			curNode.path[v].endCount++
		}
		curNode = curNode.path[v]
	}
}

func (tree *PrefixTree) Search(str string) int {
	if str == "" {
		return tree.root.endCount
	}

	cArr := strings.Split(str, "")
	curNode := tree.root

	for _, v := range cArr {
		if curNode == nil {
			break
		}
		curNode = curNode.path[v]
	}

	if curNode == nil {
		return 0
	}
	return curNode.endCount
}

func (tree *PrefixTree) PrefixCount(str string) int {
	if str == "" {
		return tree.root.passCount
	}

	cArr := strings.Split(str, "")
	curNode := tree.root

	for _, v := range cArr {
		if curNode == nil {
			break
		}
		curNode = curNode.path[v]
	}

	if curNode == nil {
		return 0
	}
	return curNode.passCount
}

func (tree *PrefixTree) Delete(str string) {
	if tree.Search(str) <= 0 {
		return
	}

	tree.root.passCount--
	if str == "" {
		tree.root.endCount--
	}

	cArr := strings.Split(str, "")
	curNode := tree.root

	for i, v := range cArr {
		curNode.path[v].passCount--
		if i == len(cArr)-1 {
			curNode.path[v].endCount--
			if curNode.path[v].endCount == 0 && curNode.path[v].passCount == 0 {
				curNode.path[v] = nil
				break
			}
		}
		curNode = curNode.path[v]
	}
}

func NewPrefixRight() *PrefixRight {
	return &PrefixRight{
		data: make(map[string]int),
	}
}

type PrefixRight struct {
	data map[string]int
}

func (rt *PrefixRight) Insert(str string) {
	if _, ok := rt.data[str]; ok {
		rt.data[str]++
	} else {
		rt.data[str] = 1
	}
}

func (rt *PrefixRight) Search(str string) int {
	return rt.data[str]
}

func (rt *PrefixRight) Delete(str string) {
	if rt.data[str] <= 1 {
		delete(rt.data, str)
	} else {
		rt.data[str]--
	}
}

func (rt *PrefixRight) PrefixCount(str string) int {
	res := 0
	for k, v := range rt.data {
		if strings.HasPrefix(k, str) {
			res += v
		}
	}
	return res
}
