package gofalsesharing

import (
	"math"
	"runtime"
	"sync"
)

var P = runtime.GOMAXPROCS(0) // P equal number of HW threads.

// Count counts the number of instances of needle in s.
func Count(s string, needle byte) int {
	cnt := 0
	for pos := 0; pos < len(s); pos++ {
		if s[pos] == needle {
			cnt++
		}
	}

	return cnt
}

// CountConcurrent concurrently counts the number of instances of needle in s.
func CountConcurrent(s string, needle byte) int {
	results := make([]int, P)

	var wg sync.WaitGroup
	wg.Add(P)

	partSize := int(math.Ceil(float64(len(s)) / float64(P)))
	for i := 0; i < P; i++ {
		i := i
		start := i * partSize
		end := min(start + partSize, len(s))
		go func() {
			for pos := start; pos < end; pos++ {
				if s[pos] == needle {
					results[i]++
				}
			}

			wg.Done()
		}()
	}

	wg.Wait()

	cnt := 0
	for _, res := range results {
		cnt += res
	}

	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}