package main

import (
	"fmt"
	"time"
)

func main() {
	// 無緩衝
	fmt.Println("No buffer channel")
	channel := make(chan int)

	go func() {
		defer fmt.Println("sub go end")

		fmt.Println("sub go start")

		channel <- 666
	}()

	num := <-channel
	fmt.Println(num)

	// 有緩衝
	fmt.Println("Buffer channel")
	bufferChannel := make(chan int, 3)

	go func() {
		defer fmt.Println("sub go1 end")

		for i := 0; i < cap(bufferChannel)+1; i++ {
			bufferChannel <- i
			fmt.Println("sub go run, ele=", i)
		}
	}()

	/* go func() {
		defer fmt.Println("sub go2 end")

		for {
			time.Sleep(1 * time.Second)
			if len(bufferChannel) > 0 {
				fmt.Println(<-bufferChannel)
			} else {
				return
			}
		}
	}() */

	time.Sleep(2 * time.Second)

	for i := 0; i < cap(bufferChannel)+1; i++ {
		num := <-bufferChannel
		fmt.Println("num =", num)
	}

	fmt.Println("main end")
}
