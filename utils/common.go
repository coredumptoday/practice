package utils

import (
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
