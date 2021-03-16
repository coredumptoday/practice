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

	m = other.GetMidOrUpMidNextNode(l)
	fmt.Println("GetMidOrUpMidNextNode: ")
	utils.PrintLinkedList(l)
	utils.PrintLinkedListPos(l, m)
}
