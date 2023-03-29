package main

import (
	"fmt"
	"math/rand"
	"time"
)

const bufferCapacity = 10

var buffer = make(chan int, bufferCapacity)

func producer(id int) {
	for {
		data := rand.Intn(20) + 1
		fmt.Printf("Producer %d produced %d\n", id, data)
		buffer <- data
		time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
	}
}

func consumer(id int) {
	for {
		select {
		case data := <-buffer:
			result := fib(data)
			fmt.Printf("Consumer %d calculated Fibonacci(%d) = %d\n", id, data, result)
		default:
			fmt.Printf("Consumer %d found no information\n", id)
		}
		time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	for i := 1; i <= 2; i++ {
		go producer(i)
	}
	for i := 1; i <= 3; i++ {
		go consumer(i)
	}
	select {}
}
