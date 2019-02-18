package gofalsesharing

import (
	"runtime"
	"sync"
)

var CPUS = runtime.GOMAXPROCS(0) // P equal number of HW threads.

func SumLinear(a []int) int {
	var cnt int
	for i := 0; i < len(a); i++ {
		cnt += a[i]
	}
	return cnt
}

func SumParallelFalseSharing(a []int) int {
	results := make([]int, CPUS)

	var wg sync.WaitGroup
	wg.Add(CPUS)

	N := len(a)
	blockSize := (N + CPUS - 1) / CPUS
	for i := 0; i < CPUS; i++ {
		i := i
		start := i * blockSize
		end := min(blockSize*(i+1), N)
		go func() {
			for pos := start; pos < end; pos++ {
				results[i] += a[i]
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

// padded must be at least 64 bytes long to lay in one cache line
// 8+56 = 64 bytes eq 1 cache line size (for x64 platforms only)
type padded struct {
	cnt int    // 1*8 = 8 bytes
	pad [7]int // 7*8 = 56 bytes
}

func SumParallelWithPadding(a []int) int {
	results := make([]padded, CPUS)

	var wg sync.WaitGroup
	wg.Add(CPUS)

	N := len(a)
	blockSize := (N + CPUS - 1) / CPUS
	for i := 0; i < CPUS; i++ {
		i := i
		start := i * blockSize
		end := min(blockSize*(i+1), N)
		go func() {
			for pos := start; pos < end; pos++ {
				results[i].cnt += a[i]
			}

			wg.Done()
		}()
	}

	wg.Wait()

	cnt := 0
	for _, res := range results {
		cnt += res.cnt
	}

	return cnt
}

func SumParallelLocalVariable(a []int) int {
	results := make([]int, CPUS)

	var wg sync.WaitGroup
	wg.Add(CPUS)

	N := len(a)
	blockSize := (N + CPUS - 1) / CPUS
	for i := 0; i < CPUS; i++ {
		i := i
		start := i * blockSize
		end := min(blockSize*(i+1), N)
		go func() {
			var cnt int
			for pos := start; pos < end; pos++ {
				cnt += a[i]
			}
			results[i] = cnt

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
