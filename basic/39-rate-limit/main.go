package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	// 限制器，限制处理的速率，每隔200ms处理一个请求
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	// 面对突发的请求，先直接处理一部分
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 剩下的部分每隔200ms处理一个请求
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 模拟突发的请求
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// 输出结果
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("burstyRequests", req, time.Now())
	}
}
