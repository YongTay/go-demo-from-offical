package main

import "fmt"

func main() {
	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("receive message", msg)
	default:
		fmt.Println("no message receive")
	}

	msg := "hi"
	select {
	// 由于message chan 不是一个buffer，所以无法送出数据
	// 因此也会走default
	case message <- msg:
		fmt.Println("receive message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-message:
		fmt.Println("receive message", msg)
	case sig := <-signals:
		fmt.Println("receive signal", sig)
	default:
		fmt.Println("no activity")
	}

}
