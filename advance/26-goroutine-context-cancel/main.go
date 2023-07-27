package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop:", name)
			return
		default:
			fmt.Printf("worker[%v] do something...\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// ！！！注意只有在goroutine才有效
	go worker(ctx, "one")
	go worker(ctx, "two")
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
