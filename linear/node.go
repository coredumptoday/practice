package linear

type Node struct {
	Value int32
	Next  *Node
}

type NodeJmp struct {
	Value int32
	Next  *NodeJmp
	Jmp   *NodeJmp
}

type DoubleNode struct {
	Value int32
	Next  *DoubleNode
	Prev  *DoubleNode
}
