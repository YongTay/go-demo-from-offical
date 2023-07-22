package main

import (
	"fmt"
	"time"
)

// ticker
// 定时循环调用多次，直到手动停止（stop）
// timer 到达时间后执行一次，也可以在未执行前停止（stop）
func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
