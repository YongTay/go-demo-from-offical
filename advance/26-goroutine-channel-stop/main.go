package main

import (
	"fmt"
	"time"
)

// 无法控制goroutine正常退出
func main() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("1 exit")
				break
			default:
				fmt.Println("1 do something...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("2 exit")
				break
			default:
				fmt.Println("2 do something...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	stop <- true
	stop <- true
	time.Sleep(3 * time.Second)
}
