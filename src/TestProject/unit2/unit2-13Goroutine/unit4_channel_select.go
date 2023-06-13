package main

import "fmt"

func fibonacci(channel, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case channel <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	channel, quit := make(chan int), make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel)
		}

		quit <- 0
	}()

	fibonacci(channel, quit)
}
