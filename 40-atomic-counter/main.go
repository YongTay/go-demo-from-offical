package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var sum uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 10000; c++ {
				atomic.AddUint64(&sum, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	println("sum", sum)
}
