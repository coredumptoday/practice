package test

import (
	"testing"
)

const NUM = 100000

func BenchmarkMapDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, NUM)
		for i := 0; i < NUM; i++ {
			m[i] = i
		}
	}

}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for i := 0; i < NUM; i++ {
			m[i] = i
		}
	}
}
