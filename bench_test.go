package gofalsesharing

import (
	"math"
	"runtime"
	"testing"
)

func in(magnitude int) []int {
	size := int(math.Pow10(magnitude))

	a := make([]int, size)
	//for i := 0; i < size; i++ {
	//	a[i] = rand.Int()
	//}

	return a
}

// avoid compiler optimisations
// @see https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
var result int

func BenchmarkSum(b *testing.B) {
	arr := in(8)

	b.Run("Linear", func(b *testing.B) {
		runtime.GC()
		b.ResetTimer()

		var r int
		for n := 0; n < b.N; n++ {
			r = SumLinear(arr)
		}

		result = r
	})

	b.Run("ParallelFalseSharing", func(b *testing.B) {
		runtime.GC()
		b.ResetTimer()

		var r int
		for n := 0; n < b.N; n++ {
			r = SumParallelFalseSharing(arr)
		}

		result = r
	})

	b.Run("ParallelWithPadding", func(b *testing.B) {
		runtime.GC()
		b.ResetTimer()

		var r int
		for n := 0; n < b.N; n++ {
			r = SumParallelWithPadding(arr)
		}

		result = r
	})

	b.Run("ParallelLocalVariable", func(b *testing.B) {
		runtime.GC()
		b.ResetTimer()

		var r int
		for n := 0; n < b.N; n++ {
			r = SumParallelLocalVariable(arr)
		}

		result = r
	})
}
