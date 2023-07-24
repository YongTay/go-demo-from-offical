package main

import (
	"fmt"
	"time"
)

func main() {
	var p = fmt.Println
	now := time.Now()
	p(now)

	// 自定义时间
	then := time.Date(2023, 5, 7, 8, 12, 33, 0, time.Local)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Date())
	p(then.Location())
	p(now.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())

	p(then.Add(diff))
}
