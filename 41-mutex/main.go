package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0, "c": 0},
	}

	doIncrement := func(name string, max int) {
		for i := 0; i < max; i++ {
			c.inc(name)
		}
	}
	var wg sync.WaitGroup
	names := []string{"a", "b", "c"}
	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			doIncrement(name, 10000)
			wg.Done()
		}(name)
	}

	wg.Wait()

	fmt.Println(c.counters)
}
