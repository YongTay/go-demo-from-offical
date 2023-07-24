package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("sent job ", i)
		time.Sleep(time.Millisecond * 100)
		jobs <- i
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
