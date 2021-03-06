package main

import (
	"fmt"
	"time"
)

var pool = 100

// Goroutine1 : 模拟生产者,打印奇数
func Goroutine1(p chan int) {
	for i := 1; i <= pool; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Printf("Goroutine1 : %d\n", i)
		}
	}
}

// Goroutine2 ： 模拟消费者，打印偶数
func Goroutine2(p chan int) {
	for i := 1; i <= pool; i++ {
		<-p
		if i%2 == 0 {
			fmt.Printf("Goroutine2 : %d\n", i)
		}
	}
}

func main() {
	msg := make(chan int)

	go Goroutine1(msg)
	go Goroutine2(msg)

	time.Sleep(time.Second)
}
