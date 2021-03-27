package test

import (
	"reflect"
	"testing"
	"unsafe"
)

func string2bytes_normal(s string) []byte {
	return []byte(s)
}
func string2bytes(s string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bs := reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bs))
}

func bytes2string_normal(bs []byte) string {
	return string(bs)
}
func bytes2string(bs []byte) string {
	bsh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	s := reflect.StringHeader{
		Data: bsh.Data,
		Len:  bsh.Len,
	}
	return *(*string)(unsafe.Pointer(&s))
}

func BenchmarkString2Bytes_normal(b *testing.B) {
	s := "hello world"
	for i := 0; i < b.N; i++ {
		_ = string2bytes_normal(s)
	}
}

func BenchmarkString2Bytes(b *testing.B) {
	s := "hello world"
	for i := 0; i < b.N; i++ {
		_ = string2bytes(s)
	}
}

func BenchmarkBytes2String_normal(b *testing.B) {
	bs := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	for i := 0; i < b.N; i++ {
		_ = bytes2string_normal(bs)
	}
}

func BenchmarkBytes2String(b *testing.B) {
	bs := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	for i := 0; i < b.N; i++ {
		_ = bytes2string(bs)
	}
}
