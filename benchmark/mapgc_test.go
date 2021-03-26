package test

/**
 * User: zhangyu
 * Email: zyzhangyu@didichuxing.com
 * Date: 2020-04-30
 * Time: 14:17
 * Desc:
 */

import (
	"testing"
)

type Value struct {
	A int64
	B int64
	C int64
	D int64
	E int64
	F int64
	G int64
	H int64
	I int64
}

type KeyWithPointer struct {
	A *uint64
	B uint64
	C uint64
	D uint64
	E uint64
	F uint64
}

type KeyWithNoPointer struct {
	A uint64
	B uint64
	C uint64
	D uint64
	E uint64
	F uint64
}

func newKeyWithPointer(a uint64) KeyWithPointer {
	oKeyWithPointer := KeyWithPointer{
		A: &a,
	}
	return oKeyWithPointer
}

func newKeyWithNoPointer(a uint64) KeyWithNoPointer {
	oKeyWithNoPointer := KeyWithNoPointer{
		A: a,
		B: a,
		C: a,
		D: a,
		E: a,
		F: a,
	}
	return oKeyWithNoPointer
}

func BenchmarkMapWithPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var aArr [NUM]KeyWithPointer
		for i := 0; i < NUM; i++ {
			aArr[i] = newKeyWithPointer(uint64(i))
		}
		m := make(map[KeyWithPointer]Value, NUM)
		for i := 0; i < NUM; i++ {
			m[aArr[i]] = Value{}
		}
		for k, _ := range m {
			delete(m, k)
		}
	}
}

func BenchmarkMapWithNoPointer(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var aArr [NUM]KeyWithNoPointer
		for i := 0; i < NUM; i++ {
			aArr[i] = newKeyWithNoPointer(uint64(i))
		}
		m := make(map[KeyWithNoPointer]Value, NUM)
		for i := 0; i < NUM; i++ {
			m[aArr[i]] = Value{}
		}
		for k, _ := range m {
			delete(m, k)
		}
	}
}
