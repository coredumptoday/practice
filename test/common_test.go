package test

import (
	"github.com/coredumptoday/practice/utils"
	"testing"
)

//获取随机数[0,1)
func TestGetRandNum(t *testing.T) {
	t.Log(utils.GetRandNum())
}

//生成随机slice
func TestGenerateRandomSlice(t *testing.T) {
	t.Log(utils.GenerateRandomSlice(50, 111))
}

//对比两个slice值是否想等
func TestEqual(t *testing.T) {
	a := utils.GenerateRandomSlice(50, 111)
	t.Log(utils.IsSliceEqual(a, a))
}
