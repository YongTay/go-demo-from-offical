package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	fmt.Println("start:", time.Now())
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop:", name, time.Now())
			return
		default:
			fmt.Println(name, "do something...", time.Now())
			options := ctx.Value("options")
			fmt.Println("options:", options)
			time.Sleep(1 * time.Second)
		}
	}
}

// WithTimeout 从启动的那一刻开始计算，到达时间后自动退出
// 本质WithDeadline(parent, time.Now().Add(timeout))
// 也可以立即手动cancel
func main() {
	options := map[string]any{
		"message": "a message",
	}
	// 一秒后自自动退出
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	ctxV := context.WithValue(ctx, "options", &options)
	go worker(ctxV, "One")
	go worker(ctxV, "Two")
	time.Sleep(time.Second * 3)
}
