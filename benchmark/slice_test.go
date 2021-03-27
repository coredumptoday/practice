package test

import "testing"

var N = 100000

func BenchmarkAppendWithoutCapacity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr []int
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkAppendWithCapLessLen100Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N/100)
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkAppendWithCapLessLen10Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N/10)
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkAppendWithCapLessLen3Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N/3)
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkAppendWithCapEqualLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N)
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkAppendWithCapGreaterLen3Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N*3)
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkAppendWithCapGreaterLen10Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N*10)
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}
func BenchmarkAppendWithCapGreaterLen100Times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N*100)
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkAppendWithoutCapacityReuse(b *testing.B) {
	var arr []int
	for i := 0; i < b.N; i++ {
		arr = arr[:0]
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkAppendWithCapEqualLenReuse(b *testing.B) {
	arr := make([]int, N)
	for i := 0; i < b.N; i++ {
		arr = arr[:0]
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkAppendWithCapGreaterLen10TimesReuse(b *testing.B) {
	arr := make([]int, N*10)
	for i := 0; i < b.N; i++ {
		arr = arr[:0]
		for j := 0; j < N; j++ {
			arr = append(arr, j)
		}
	}
}
