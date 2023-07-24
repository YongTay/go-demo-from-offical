package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	select {
	case msg := <-c1:
		fmt.Println("receive:", msg)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout handler...")
	}
	fmt.Println("block?") // yes

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "two"
	}()

	select {
	case msg := <-c2:
		fmt.Println("receive:", msg)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout handler")
	}
}
