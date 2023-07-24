package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 为什么需要循环，因为一次只能处理一个channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("receive:", msg1)
		case msg2 := <-c2:
			fmt.Println("receive:", msg2)
		}
	}
}
