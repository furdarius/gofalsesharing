package gofalsesharing

import (
	"math"
	"math/rand"
	"runtime"
	"testing"
)

var result1, result2 int

func in(magnitude int) (s string, needle byte) {
	size := int(math.Pow10(magnitude))
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	s = string(buf)
	needle = s[rand.Intn(size-1)]

	return s, needle
}

var s, needle = in(8)

func BenchmarkCountLinear(b *testing.B) {
	var r int

	runtime.GC()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r = Count(s, needle)
	}

	result1 = r
}

func BenchmarkCountConcurrent(b *testing.B) {
	var r int

	runtime.GC()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r = CountConcurrent(s, needle)
	}

	result2 = r
}
