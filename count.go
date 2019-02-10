package gofalsesharing

// Count counts the number of instances of needle in s.
func Count(s string, needle rune) int {
	cnt := 0
	for _, ch := range s {
		if ch == needle {
			cnt++
		}
	}
	return cnt
}
//
//// CountConcurrent concurrently counts the number of instances of needle in s.
//func CountConcurrent(s string, needle rune) int {
//	P := runtime.GOMAXPROCS(0) // P equal number of HW threads.
//
//	results := make([]int, P)
//
//	var wg sync.WaitGroup
//	wg.Add(P)
//
//	partSize := int(math.Ceil(float64(len(s)) / float64(P)))
//	for i := 0; i < P; i++ {
//		i := i
//		start := i * partSize
//		end := min(start + partSize, len(s))
//		go func() {
//			for _, chr := range s[start:end] {
//				if chr == needle {
//					results[i]++
//				}
//			}
//
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//
//	cnt := 0
//	for _, res := range results {
//		cnt += res
//	}
//
//	return cnt
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//
//	return b
//}