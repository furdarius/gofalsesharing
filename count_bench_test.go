package gofalsesharing
//
//import (
//	"math"
//	"math/rand"
//	"runtime"
//	"strings"
//	"testing"
//)
//
//// avoid compiler optimisations
//// @see https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
//var result1, result2, result3, result4 int
//
//func in(magnitude int) []int {
//	size := int(math.Pow10(magnitude))
//
//	a := make([]int, size)
//	for i := 0; i < size; i++ {
//		a[i] = rand.Int()
//	}
//
//	return a
//}
//
//// inRep returns string consisting of a repeating character.
//func inRep(magnitude int) (s string, needle byte) {
//	size := int(math.Pow10(magnitude))
//	s = strings.Repeat("a", size)
//	needle = s[rand.Intn(size-1)]
//
//	return s, needle
//}
//
//var s, needle = inRep(6)
//
//func BenchmarkCountLinear(b *testing.B) {
//	var r int
//
//	runtime.GC()
//	b.ResetTimer()
//
//	for n := 0; n < b.N; n++ {
//		r = CountLinear(s, needle)
//	}
//
//	result1 = r
//}
//
//func BenchmarkCountConcurrentFalseSharing(b *testing.B) {
//	var r int
//
//	runtime.GC()
//	b.ResetTimer()
//
//	for n := 0; n < b.N; n++ {
//		r = CountConcurrentFalseSharing(s, needle)
//	}
//
//	result2 = r
//}
//
//func BenchmarkCountConcurrentWithPadding(b *testing.B) {
//	var r int
//
//	runtime.GC()
//	b.ResetTimer()
//
//	for n := 0; n < b.N; n++ {
//		r = CountConcurrentWithPadding(s, needle)
//	}
//
//	result3 = r
//}
//
//func BenchmarkCountConcurrentWithLocalVar(b *testing.B) {
//	var r int
//
//	runtime.GC()
//	b.ResetTimer()
//
//	for n := 0; n < b.N; n++ {
//		r = CountConcurrentWithLocalVar(s, needle)
//	}
//
//	result3 = r
//}
