package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Println("id", id, "start work")
	time.Sleep(time.Second)
	fmt.Println("id", id, "finished work")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		// i := i
		// go func() {
		// 	defer wg.Done()
		// 	worker(i)
		// }()
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}

	wg.Wait()
}
