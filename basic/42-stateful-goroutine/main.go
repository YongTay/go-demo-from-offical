package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

// channel的读取会阻塞，所以可以用for循环去实现一直监听
func main() {

	// 用于计数
	var readOps uint64
	var writeOps uint64

	// 读写的channel
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 一直等待监听channel
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key] // 读取通道数据，读取完成后会设置标志
			case write := <-writes:
				state[write.key] = write.val // 保存写通道的数据，在读通道会进行读取
				write.resp <- true           // 设置写完标志
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read                 // 往读通道写数据
				<-read.resp                   // 等待读通道读取完
				atomic.AddUint64(&readOps, 1) // 计数
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write                // 往写通道里写数据
				<-write.resp                   // 等待写通道写入完成
				atomic.AddUint64(&writeOps, 1) // 计数
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 阻塞等待时间到，系统将继续运行，然后退出
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOpsFinal", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOpsFinal", writeOpsFinal)
}
