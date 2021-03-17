package other

import "github.com/coredumptoday/practice/linear"

//奇数节点获取中点，偶数节点获取上中点（偏向前半）
func GetMidOrUpMidNode(head *linear.Node) *linear.Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

//奇数节点获取中点，偶数节点获取下中点（偏向后半）
func GetMidOrDownMidNode(head *linear.Node) *linear.Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	slow := head.Next
	fast := head.Next

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

//奇数节点获取中点，偶数节点获取上中点前一个节点
func GetMidOrUpMidPreNode(head *linear.Node) *linear.Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	slow := head
	fast := head.Next.Next

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

//奇数节点获取中点，偶数节点获取下中点前一个节点
func GetMidOrDownMidPreNode(head *linear.Node) *linear.Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	slow := head
	fast := head.Next

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
