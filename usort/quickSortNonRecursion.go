package usort

import (
	"container/list"
	"github.com/coredumptoday/practice/utils"
)

type pos struct {
	l int
	r int
}

func genRange(arr []int, l, r int) []int {
	if l > r {
		return []int{-1, -1}
	}
	if l == r {
		return []int{l, r}
	}

	less := l - 1
	more := r
	idx := l

	for idx < more {
		if arr[idx] == arr[r] {
			idx++
		} else if arr[idx] < arr[r] {
			utils.IntSliceSwap(arr, less+1, idx)
			less++
			idx++
		} else {
			utils.IntSliceSwap(arr, more-1, idx)
			more--
		}
	}

	utils.IntSliceSwap(arr, more, r)
	return []int{less + 1, more}
}

func processNonRecursion(arr []int, l, r int) {
	if l >= r {
		return
	}
	x := int(utils.GetRandNum() * float32(r-l+1))
	utils.IntSliceSwap(arr, l+x, r)

	stack := list.New()

	m := genRange(arr, l, r)
	stack.PushFront(pos{
		l: l,
		r: m[0] - 1,
	})
	stack.PushFront(pos{
		l: m[0] + 1,
		r: r,
	})

	for stack.Len() > 0 {
		top := stack.Front()
		stack.Remove(top)

		l := top.Value.(pos).l
		r := top.Value.(pos).r

		if l >= r {
			continue
		}
		x := int(utils.GetRandNum() * float32(r-l+1))
		utils.IntSliceSwap(arr, l+x, r)

		m := genRange(arr, l, r)
		stack.PushFront(pos{
			l: l,
			r: m[0] - 1,
		})
		stack.PushFront(pos{
			l: m[0] + 1,
			r: r,
		})
	}
}

func QuickSortNonRecursion(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	processNonRecursion(arr, 0, len(arr)-1)
}
