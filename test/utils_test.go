package test

import (
	"github.com/coredumptoday/practice/utils"
	"testing"
)

func TestGetRandNum(t *testing.T)  {
	t.Log(utils.GetRandNum())
}

func TestGenerateRandomSlice(t *testing.T)  {
	t.Log(utils.GenerateRandomSlice(50, 111))
}

func TestEqual(t *testing.T)  {
	a := utils.GenerateRandomSlice(50, 111)
	t.Log(utils.IsSliceEqual(a, a))
}