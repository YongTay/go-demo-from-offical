package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer1 was fired")

	time2 := time.NewTimer(time.Second)

	go func() {
		<-time2.C
		fmt.Println("timer2 was fired")
	}()
	stop := time2.Stop()
	if stop {
		fmt.Println("timer2 was stopped")
	}

	time.Sleep(2 * time.Second)
}
