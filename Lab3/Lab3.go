package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	buffer = make(chan int)
)

func producer(id int) {
	for {

		data := rand.Intn(20) + 1    // Generate a random integer between 1 and 20
		duration := rand.Intn(4) + 1 // Generate a random integer between 1 and 4
		buffer <- data               // Add the value to the buffer
		fmt.Printf("Producer %d produced data: %d TimeSleep %d\n", id, data, duration)

		// Sleep for a random duration between 1 and 4 seconds
		time.Sleep(time.Duration(duration) * time.Second)
	}
}

func consumer(id int) {
	for {
		data, ok := <-buffer         // Add the value to the buffer
		duration := rand.Intn(4) + 1 // Generate a random integer between 1 and 4
		if ok {
			fib := fib(data) // Calculate the Fibonacci
			fmt.Printf("Consumer %d Fibonacci value for %d: %d TimeSleep %d\n", id, data, fib, duration)
		} else {
			fmt.Println("No data in buffer")
		}
		// Sleep for a random duration between 1 and 4 seconds
		time.Sleep(time.Duration(duration) * time.Second)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// Create the producer threads
	for i := 1; i <= 2; i++ {
		go producer(i)
	}
	// Create the consumer threads
	for i := 1; i <= 3; i++ {
		go consumer(i)
	}

	fmt.Scanln()
}
