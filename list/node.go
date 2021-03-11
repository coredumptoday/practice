package list

type Node struct {
	Value int32
	Next  *Node
}

type DoubleNode struct {
	Value int32
	Next  *DoubleNode
	Prev  *DoubleNode
}
