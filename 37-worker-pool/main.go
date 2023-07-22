package main

import (
	"fmt"
	"time"
)

// 执行任务函数
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("id", id, "start job", j)
		time.Sleep(2 * time.Second)
		fmt.Println("id", id, "job finished", j)
		results <- j * 2
	}
}

/**
* 任务保存在jobs中，worker函数会一直取出函数直到结束
 */

func main() {
	jobNumber := 5
	jobs := make(chan int, jobNumber)
	results := make(chan int, jobNumber)
	// 最后关闭
	defer close(jobs)
	defer close(results)

	// 创建协程去执行任务
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// 创建任务
	for i := 0; i < jobNumber; i++ {
		jobs <- i
	}

	// 消费任务
	for i := 0; i < jobNumber; i++ {
		<-results
	}
}
