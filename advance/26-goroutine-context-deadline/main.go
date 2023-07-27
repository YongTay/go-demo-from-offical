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
			fmt.Println("stop", name)
			return
		default:
			fmt.Println("do something", name)
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	// 从启动的这一刻开始计算时间，到时间后就自动退出
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	ctxV := context.WithValue(ctx, "message", "deadline")
	go worker(ctxV, "One")
	go worker(ctxV, "Two")
	time.Sleep(5 * time.Second)
}
