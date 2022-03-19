package testing

import "testing"

var smallSize = 100
var bigSize = 1000000

func BenchmarkMakeSmallSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]byte, smallSize)
	}
}

func BenchmarkMakeSmallMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[int]bool, smallSize)
	}
}

func BenchmarkMakeSmallChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(chan int, smallSize)
	}
}

func BenchmarkMakeBigSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]byte, bigSize)
	}
}

func BenchmarkMakeBigMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[int]bool, bigSize)
	}
}

func BenchmarkMakeBigChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(chan int, bigSize)
	}
}

type SmallObj struct {
	a string
	b string
	c string
}

type BigObj struct {
	a [100000]string
	b [100000]string
	c [100000]string
}

func BenchmarkNewSmallObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = new(SmallObj)
	}
}

func BenchmarkNewBigObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = new(BigObj)
	}
}
