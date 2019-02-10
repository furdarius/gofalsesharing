package gofalsesharing

import (
	"math"
	"math/rand"
	"strconv"
	"testing"
)

var	result1, result2 int

func benchmarkCount(b *testing.B, s string, needle rune) {
	var r int

	for n := 0; n < b.N; n++ {
		r = Count(s, needle)
	}

	// avoid compiler optimisations
	// @see: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
	result1 = r
}

func benchmarkCountConcurrent(b *testing.B, s string, needle rune) {
	var r int

	for n := 0; n < b.N; n++ {
		r = CountConcurrent(s, needle)
	}

	// avoid compiler optimisations
	// @see: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
	result2 = r
}

func in(magnitude int) (s string, needle rune) {
	size := int(math.Pow10(magnitude))
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	s = string(buf)
	needle = rune(s[rand.Intn(size-1)])

	return s, needle
}

func BenchmarkCount(b *testing.B) {
	for magnitude := 6; magnitude <= 8; magnitude++ {
		s, needle := in(magnitude)

		b.Run("Magnitude"+strconv.Itoa(magnitude), func(b *testing.B) {
			benchmarkCount(b, s, needle)
		})
	}
}

func BenchmarkCountConcurrent(b *testing.B) {
	for magnitude := 6; magnitude <= 8; magnitude++ {
		s, needle := in(magnitude)

		b.Run("Magnitude"+strconv.Itoa(magnitude), func(b *testing.B) {
			benchmarkCountConcurrent(b, s, needle)
		})
	}
}


