package utils

import (
	"github.com/coredumptoday/practice/tree"
	"strings"
)

func GenRandomString(charArr []string, strLen int) string {
	ans := make([]string, int(GetRandNum()*float32(strLen))+1)

	for i := range ans {
		val := int(GetRandNum() * float32(26))
		ans[i] = charArr[val]
	}

	return strings.Join(ans, "")
}

func GenerateRandomStringArray(arrLen, strLen int) []string {
	ans := make([]string, int(GetRandNum()*float32(arrLen))+1)
	for i := range ans {
		ans[i] = GenRandomString(tree.CharArr, strLen)
	}
	return ans
}
