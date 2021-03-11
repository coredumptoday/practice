package utils

import (
	"fmt"
	"github.com/coredumptoday/practice/list"
	"math"
	"math/rand"
	"time"
)

func GetRandNum() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}

func GetValue(maxValue int) int {
	return int(float32(maxValue+1)*GetRandNum()) - int(float32(maxValue)*GetRandNum())
}

func GenerateRandomSlice(maxSize, maxValue int) []int {
	cap := int(float32(maxSize+1) * GetRandNum())
	arr := make([]int, cap)

	for i, _ := range arr {
		arr[i] = GetValue(maxValue)
	}

	return arr
}

func IntSliceSwap(arr []int, i, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}

func IntInSlice(arr []int, v int) bool {
	if arr == nil || len(arr) <= 0 {
		return false
	}

	for _, val := range arr {
		if val == v {
			return true
		}
	}

	return false
}

func IsSliceEqual(a1, a2 []int) bool {
	if (a1 == nil && a2 != nil) || (a1 != nil && a2 == nil) {
		return false
	}

	if a1 == nil && a2 == nil {
		return true
	}

	if len(a1) != len(a2) {
		return false
	}

	for i, v1 := range a1 {
		if a2[i] != v1 {
			return false
		}
	}

	return true
}

func mkRandomNum(a int) int {
	return int((float32(a)*GetRandNum() + 1) - (float32(a)*GetRandNum() + 1))
}

func MKRandomSlice(maxKinds, a, k, m int) ([]int32, int32) {
	ktimeNum := mkRandomNum(a)

	// 真命天子出现的次数
	times := 0
	if GetRandNum() < 0.5 {
		times = k
	} else {
		times = int(GetRandNum()*float32(m-1) + 1)
	}

	// 2
	numKinds := int(GetRandNum()*float32(maxKinds) + 2)

	// k * 1 + (numKinds - 1) * m
	arr := make([]int32, times+(numKinds-1)*m)

	index := 0
	for ; index < times; index++ {
		arr[index] = int32(ktimeNum)
	}

	set := make(map[int]bool, numKinds)
	set[ktimeNum] = true
	numKinds--

	for numKinds != 0 {
		curNum := 0
		for {
			curNum = mkRandomNum(a)
			if !set[curNum] {
				break
			}
		}
		set[curNum] = true
		numKinds--
		for i := 0; i < m; i++ {
			arr[index] = int32(curNum)
			index++
		}
	}

	for i := 0; i < len(arr); i++ {
		j := int(GetRandNum() * float32(len(arr)))
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	}

	return arr, int32(ktimeNum)
}

func KmTest(arr []int32, k int32) int32 {
	mm := make(map[int32]int32, len(arr))

	for _, v := range arr {
		if n, ok := mm[v]; ok {
			mm[v] = n + 1
		} else {
			mm[v] = 1
		}
	}

	for i, v := range mm {
		if v == k {
			return i
		}
	}

	return math.MinInt32
}

func PrintLinkedList(head *list.Node) {
	for head != nil {
		fmt.Print(head.Value, " -> ")
		head = head.Next
	}
	fmt.Println("null")
}

func PrintDoubleList(head *list.DoubleNode) {
	for head != nil {
		if head.Prev == nil {
			fmt.Print("null <-")
		} else {
			fmt.Print("<-")
		}
		fmt.Print(head.Value)
		if head.Next == nil {
			fmt.Print("-> null")
		} else {
			fmt.Print("->")
		}
		head = head.Next
	}
	fmt.Println()
}

func GenerateRandomLinkedList(length, value int32) *list.Node {
	size := int32(GetRandNum() * float32(length+1))
	if size == 0 {
		return nil
	}

	head := list.Node{
		Value: int32(GetRandNum() * float32(value+1)),
		Next:  nil,
	}
	prev := &head
	size--

	for size != 0 {
		cur := list.Node{
			Value: int32(GetRandNum() * float32(value+1)),
			Next:  nil,
		}
		(*prev).Next = &cur
		prev = &cur
		size--
	}

	return &head
}

func GenerateRandomDoubleList(length, value int32) *list.DoubleNode {
	size := int32(GetRandNum() * float32(length+1))
	if size == 0 {
		return nil
	}

	head := list.DoubleNode{
		Value: int32(GetRandNum() * float32(value+1)),
		Next:  nil,
		Prev:  nil,
	}
	prev := &head
	size--

	for size != 0 {
		cur := list.DoubleNode{
			Value: int32(GetRandNum() * float32(value+1)),
			Next:  nil,
			Prev:  prev,
		}
		(*prev).Next = &cur
		prev = &cur
		size--
	}

	return &head
}

func GetLinkedListSlice(head *list.Node) []int32 {
	l := make([]int32, 0)
	for head != nil {
		l = append(l, head.Value)
		head = head.Next
	}
	return l
}

func GetDoubleListSlice(head *list.DoubleNode) []int32 {
	l := make([]int32, 0)
	for head != nil {
		l = append(l, head.Value)
		head = head.Next
	}
	return l
}

func PrintSliceRev(a []int32, isDouble bool) {
	t := ""
	if isDouble {
		t = " <-> "
	} else {
		t = " -> "
	}

	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(a[i], t)
	}
	fmt.Println("null")
}

func CheckSliceAndLinkedList(a []int32, head *list.Node) bool {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != head.Value {
			return false
		}
		head = head.Next
	}
	return true
}

func CheckSliceAndDoubleList(a []int32, head *list.DoubleNode) bool {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != head.Value {
			return false
		}
		head = head.Next
	}
	return true
}
