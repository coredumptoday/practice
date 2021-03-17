package tree

type data struct {
	isBST bool
	max   int
	min   int
}

func process(cur *BinTreeNode) *data {
	if cur == nil {
		return nil
	}

	leftInfo := process(cur.LeftNode)
	rightInfo := process(cur.RightNode)

	isBst := true
	max := cur.Value
	min := cur.Value
	if leftInfo != nil {
		isBst = isBst && leftInfo.isBST
		if cur.Value <= leftInfo.max {
			isBst = false
		}
		if max < leftInfo.max {
			max = leftInfo.max
		}
		if min > leftInfo.min {
			min = leftInfo.min
		}
	}
	if rightInfo != nil {
		isBst = isBst && rightInfo.isBST

		if cur.Value >= rightInfo.min {
			isBst = false
		}
		if max < rightInfo.max {
			max = rightInfo.max
		}
		if min > rightInfo.min {
			min = rightInfo.min
		}
	}

	return &data{
		isBST: isBst,
		max:   max,
		min:   min,
	}
}

func IsBinSearchTree(cur *BinTreeNode) bool {
	if cur == nil {
		return true
	}
	d := process(cur)
	return d.isBST
}
