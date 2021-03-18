package test

import (
	"fmt"
	"github.com/coredumptoday/practice/linear"
	"github.com/coredumptoday/practice/other"
	"github.com/coredumptoday/practice/utils"
	"sort"
	"testing"
)

func TestLessSum(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		a1 := utils.GenerateRandomSlice(maxSize, maxValue)
		a2 := make([]int, len(a1))
		copy(a2, a1)

		s1 := other.LessSum(a1)
		s2 := other.Comparator(a2)

		if s1 != s2 {
			succeed = false
			fmt.Println("sum1: ", s1, "a1: ", a1)
			fmt.Println("sum2: ", s2, "a2: ", a2)
			t.Fail()
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("lessSum succeed count: ", i+1)
			fmt.Println("sum1: ", s1, "a1: ", a1)
			fmt.Println("sum2: ", s2, "a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fail!")
	}
}

func TestRevCount(t *testing.T) {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		a1 := utils.GenerateRandomSlice(maxSize, maxValue)
		a2 := make([]int, len(a1))
		copy(a2, a1)

		s1 := other.ReversedCount(a1)
		s2 := other.RevComparator(a2)

		if s1 != s2 {
			succeed = false
			fmt.Println("sum1: ", s1, "a1: ", a1)
			fmt.Println("sum2: ", s2, "a2: ", a2)
			t.Fail()
			break
		}

		if (i+1)%20000 == 0 {
			fmt.Println("RevCount succeed count: ", i+1)
			fmt.Println("sum1: ", s1, "a1: ", a1)
			fmt.Println("sum2: ", s2, "a2: ", a2)
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fail!")
	}
}

func TestComparator(t *testing.T) {
	student1 := &other.Student{
		Id:   4,
		Age:  40,
		Name: "A",
	}
	student2 := &other.Student{
		Id:   4,
		Age:  21,
		Name: "B",
	}
	student3 := &other.Student{
		Id:   3,
		Age:  12,
		Name: "C",
	}
	student4 := &other.Student{
		Id:   3,
		Age:  62,
		Name: "D",
	}
	student5 := &other.Student{
		Id:   3,
		Age:  42,
		Name: "E",
	}
	// D E C A B
	students := other.IdUpAgeDown{
		student1,
		student2,
		student3,
		student4,
		student5,
	}

	sort.Sort(students)

	for _, v := range students {
		fmt.Println("Id: ", v.Id, "Age: ", v.Age, "Name: ", v.Name)
	}
}

func TestOddEvenSort(t *testing.T) {
	maxSize := 30
	maxValue := 100

	a1 := utils.GenerateRandomSlice(maxSize, maxValue)
	fmt.Println(a1)
	idx := other.OddEvenSort(a1, 0, len(a1)-1)

	fmt.Print("[")
	for i, v := range a1 {
		fmt.Print(v)
		if i == idx {
			fmt.Print("|")
		}
		fmt.Print(" ")
	}
	fmt.Println("]")
}

func TestLinkListMid(t *testing.T) {
	l := &linear.Node{}
	l.Next = &linear.Node{Value: 1, Next: nil}
	l.Next.Next = &linear.Node{Value: 2, Next: nil}
	l.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	l.Next.Next.Next.Next = &linear.Node{Value: 4, Next: nil}
	l.Next.Next.Next.Next.Next = &linear.Node{Value: 5, Next: nil}
	l.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 6, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 7, Next: nil}
	l.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 8, Next: nil}

	m := other.GetMidOrUpMidNode(l)
	fmt.Println("GetMidOrUpMidNode: ")
	utils.PrintLinkedList(l)
	utils.PrintLinkedListPos(l, m)

	m = other.GetMidOrDownMidNode(l)
	fmt.Println("GetMidOrDownMidNode: ")
	utils.PrintLinkedList(l)
	utils.PrintLinkedListPos(l, m)

	m = other.GetMidOrUpMidPreNode(l)
	fmt.Println("GetMidOrUpMidPreNode: ")
	utils.PrintLinkedList(l)
	utils.PrintLinkedListPos(l, m)

	m = other.GetMidOrDownMidPreNode(l)
	fmt.Println("GetMidOrUpMidNextNode: ")
	utils.PrintLinkedList(l)
	utils.PrintLinkedListPos(l, m)
}

func TestIsPalindromeList(t *testing.T) {
	var head *linear.Node

	head = nil
	utils.PrintLinkedList(head)
	fmt.Println("IsPalindromeWithNSpace\t", other.IsPalindromeWithNSpace(head))
	fmt.Println("IsPalindromeWithHalfNSpace\t", other.IsPalindromeWithHalfNSpace(head))
	fmt.Println("IsPalindromeWith1Space\t", other.IsPalindromeWith1Space(head))
	utils.PrintLinkedList(head)
	fmt.Println()

	head = &linear.Node{Value: 1, Next: nil}
	utils.PrintLinkedList(head)
	fmt.Println("IsPalindromeWithNSpace\t", other.IsPalindromeWithNSpace(head))
	fmt.Println("IsPalindromeWithHalfNSpace\t", other.IsPalindromeWithHalfNSpace(head))
	fmt.Println("IsPalindromeWith1Space\t", other.IsPalindromeWith1Space(head))
	utils.PrintLinkedList(head)
	fmt.Println()

	head = &linear.Node{Value: 1, Next: nil}
	head.Next = &linear.Node{Value: 2, Next: nil}
	utils.PrintLinkedList(head)
	fmt.Println("IsPalindromeWithNSpace\t", other.IsPalindromeWithNSpace(head))
	fmt.Println("IsPalindromeWithHalfNSpace\t", other.IsPalindromeWithHalfNSpace(head))
	fmt.Println("IsPalindromeWith1Space\t", other.IsPalindromeWith1Space(head))
	utils.PrintLinkedList(head)
	fmt.Println()

	head = &linear.Node{Value: 1, Next: nil}
	head.Next = &linear.Node{Value: 1, Next: nil}
	utils.PrintLinkedList(head)
	fmt.Println("IsPalindromeWithNSpace\t", other.IsPalindromeWithNSpace(head))
	fmt.Println("IsPalindromeWithHalfNSpace\t", other.IsPalindromeWithHalfNSpace(head))
	fmt.Println("IsPalindromeWith1Space\t", other.IsPalindromeWith1Space(head))
	utils.PrintLinkedList(head)
	fmt.Println()

	head = &linear.Node{Value: 1, Next: nil}
	head.Next = &linear.Node{Value: 2, Next: nil}
	head.Next.Next = &linear.Node{Value: 3, Next: nil}
	utils.PrintLinkedList(head)
	fmt.Println("IsPalindromeWithNSpace\t", other.IsPalindromeWithNSpace(head))
	fmt.Println("IsPalindromeWithHalfNSpace\t", other.IsPalindromeWithHalfNSpace(head))
	fmt.Println("IsPalindromeWith1Space\t", other.IsPalindromeWith1Space(head))
	utils.PrintLinkedList(head)
	fmt.Println()

	head = &linear.Node{Value: 1, Next: nil}
	head.Next = &linear.Node{Value: 2, Next: nil}
	head.Next.Next = &linear.Node{Value: 1, Next: nil}
	utils.PrintLinkedList(head)
	fmt.Println("IsPalindromeWithNSpace\t", other.IsPalindromeWithNSpace(head))
	fmt.Println("IsPalindromeWithHalfNSpace\t", other.IsPalindromeWithHalfNSpace(head))
	fmt.Println("IsPalindromeWith1Space\t", other.IsPalindromeWith1Space(head))
	utils.PrintLinkedList(head)
	fmt.Println()

	head = &linear.Node{Value: 1, Next: nil}
	head.Next = &linear.Node{Value: 2, Next: nil}
	head.Next.Next = &linear.Node{Value: 3, Next: nil}
	head.Next.Next.Next = &linear.Node{Value: 1, Next: nil}
	utils.PrintLinkedList(head)
	fmt.Println("IsPalindromeWithNSpace\t", other.IsPalindromeWithNSpace(head))
	fmt.Println("IsPalindromeWithHalfNSpace\t", other.IsPalindromeWithHalfNSpace(head))
	fmt.Println("IsPalindromeWith1Space\t", other.IsPalindromeWith1Space(head))
	utils.PrintLinkedList(head)
	fmt.Println()

	head = &linear.Node{Value: 1, Next: nil}
	head.Next = &linear.Node{Value: 2, Next: nil}
	head.Next.Next = &linear.Node{Value: 2, Next: nil}
	head.Next.Next.Next = &linear.Node{Value: 1, Next: nil}
	utils.PrintLinkedList(head)
	fmt.Println("IsPalindromeWithNSpace\t", other.IsPalindromeWithNSpace(head))
	fmt.Println("IsPalindromeWithHalfNSpace\t", other.IsPalindromeWithHalfNSpace(head))
	fmt.Println("IsPalindromeWith1Space\t", other.IsPalindromeWith1Space(head))
	utils.PrintLinkedList(head)
	fmt.Println()

	head = &linear.Node{Value: 1, Next: nil}
	head.Next = &linear.Node{Value: 2, Next: nil}
	head.Next.Next = &linear.Node{Value: 3, Next: nil}
	head.Next.Next.Next = &linear.Node{Value: 2, Next: nil}
	head.Next.Next.Next.Next = &linear.Node{Value: 1, Next: nil}
	utils.PrintLinkedList(head)
	fmt.Println("IsPalindromeWithNSpace\t", other.IsPalindromeWithNSpace(head))
	fmt.Println("IsPalindromeWithHalfNSpace\t", other.IsPalindromeWithHalfNSpace(head))
	fmt.Println("IsPalindromeWith1Space\t", other.IsPalindromeWith1Space(head))
	utils.PrintLinkedList(head)
	fmt.Println()

}

func TestRangeList(t *testing.T) {
	head := &linear.Node{Value: 7, Next: nil}
	head.Next = &linear.Node{Value: 9, Next: nil}
	head.Next.Next = &linear.Node{Value: 1, Next: nil}
	head.Next.Next.Next = &linear.Node{Value: 8, Next: nil}
	head.Next.Next.Next.Next = &linear.Node{Value: 5, Next: nil}
	head.Next.Next.Next.Next.Next = &linear.Node{Value: 2, Next: nil}
	head.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 5, Next: nil}

	utils.PrintLinkedList(head)
	head = other.RangeList(head, 5)
	utils.PrintLinkedList(head)
}

func TestCopyLinkListWithRandPtr(t *testing.T) {
	var origin *linear.NodeJmp

	origin = nil
	newH := other.CopyLinkListWithRandPtr(origin)
	fmt.Println("----------- origin1:")
	utils.PrintRandLinkedList(origin)
	fmt.Println("----------- new1:")
	utils.PrintRandLinkedList(newH)

	origin = &linear.NodeJmp{Value: 1, Next: nil}
	origin.Next = &linear.NodeJmp{Value: 2, Next: nil}
	origin.Next.Next = &linear.NodeJmp{Value: 3, Next: nil}
	origin.Next.Next.Next = &linear.NodeJmp{Value: 4, Next: nil}
	origin.Next.Next.Next.Next = &linear.NodeJmp{Value: 5, Next: nil}
	origin.Next.Next.Next.Next.Next = &linear.NodeJmp{Value: 6, Next: nil}

	origin.Jmp = origin.Next.Next.Next.Next.Next                //1 -> 6
	origin.Next.Jmp = origin.Next.Next.Next.Next.Next           //2 -> 6
	origin.Next.Next.Jmp = origin.Next.Next.Next.Next           //3 -> 5
	origin.Next.Next.Next.Jmp = origin.Next.Next                //4 -> 3
	origin.Next.Next.Next.Next.Jmp = nil                        //5 -> nil
	origin.Next.Next.Next.Next.Next.Jmp = origin.Next.Next.Next //6 -> 4

	newH = other.CopyLinkListWithRandPtr(origin)
	fmt.Println("----------- origin2:")
	utils.PrintRandLinkedList(origin)
	fmt.Println("----------- new2:")
	utils.PrintRandLinkedList(newH)
}

func TestLinkListLoop(t *testing.T) {
	//0->1->2->3->4->5->6->7->8->nil
	head1 := &linear.Node{}
	head1.Next = &linear.Node{Value: 1, Next: nil}
	head1.Next.Next = &linear.Node{Value: 2, Next: nil}
	head1.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	head1.Next.Next.Next.Next = &linear.Node{Value: 4, Next: nil}
	head1.Next.Next.Next.Next.Next = &linear.Node{Value: 5, Next: nil}
	head1.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 6, Next: nil}
	head1.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 7, Next: nil}
	head1.Next.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 8, Next: nil}

	//0->1->2->3->4->5->6->7->3
	head2 := &linear.Node{}
	head2.Next = &linear.Node{Value: 1, Next: nil}
	head2.Next.Next = &linear.Node{Value: 2, Next: nil}
	head2.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	head2.Next.Next.Next.Next = &linear.Node{Value: 4, Next: nil}
	head2.Next.Next.Next.Next.Next = &linear.Node{Value: 5, Next: nil}
	head2.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 6, Next: nil}
	head2.Next.Next.Next.Next.Next.Next.Next = &linear.Node{Value: 7, Next: nil}
	head2.Next.Next.Next.Next.Next.Next.Next.Next = head2.Next.Next.Next

	//0->1->2->3->4->5->6->7->8->nil
	head3 := &linear.Node{}
	head3.Next = &linear.Node{Value: 1, Next: nil}
	head3.Next.Next = &linear.Node{Value: 2, Next: nil}
	head3.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	head3.Next.Next.Next.Next = &linear.Node{Value: 4, Next: nil}
	head3.Next.Next.Next.Next.Next = head1.Next.Next.Next.Next.Next

	//0->1->2->3->4->5->6->7->3
	head4 := &linear.Node{}
	head4.Next = &linear.Node{Value: 1, Next: nil}
	head4.Next.Next = &linear.Node{Value: 2, Next: nil}
	head4.Next.Next.Next = &linear.Node{Value: 3, Next: nil}
	head4.Next.Next.Next.Next = &linear.Node{Value: 4, Next: nil}
	head4.Next.Next.Next.Next.Next = &linear.Node{Value: 5, Next: nil}
	head4.Next.Next.Next.Next.Next.Next = head2.Next.Next

	fmt.Println("loop node:")
	fmt.Println(linear.GetLoopNode(head1))
	fmt.Println(linear.GetLoopNode(head2))

	fmt.Println("cross node no loop:")
	fmt.Println(linear.GetLinkListCrossPointNoLoop(head1, head3))

	fmt.Println("cross node with loop:")
	fmt.Println(linear.GetLinkListCrossPointWithLoop(head2, head4, linear.GetLoopNode(head2), linear.GetLoopNode(head4)))
}

func TestPrintArray(t *testing.T) {
	arr := [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
	}

	for _, v := range arr {
		for _, vv := range v {
			fmt.Print(vv, "\t")
		}
		fmt.Println()
	}

	row := len(arr)
	col := len(arr[0])
	pointCount := row * col
	times := 0
	for n := 0; ; n++ {
		for j := n; j < col-n; j++ {
			fmt.Print(arr[n][j], "\t")
			times++
		}
		if times == pointCount {
			break
		}
		for j := n + 1; j < row-n; j++ {
			fmt.Print(arr[j][col-n-1], "\t")
			times++
		}
		if times == pointCount {
			break
		}
		for j := col - n - 2; j >= n; j-- {
			fmt.Print(arr[row-n-1][j], "\t")
			times++
		}
		if times == pointCount {
			break
		}
		for j := row - n - 2; j > n; j-- {
			fmt.Print(arr[j][n], "\t")
			times++
		}
		if times == pointCount {
			break
		}
	}
	fmt.Println()
}
