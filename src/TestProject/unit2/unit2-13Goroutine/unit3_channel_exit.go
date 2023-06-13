package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- i
		}
		close(channel)
	}()

	/* for {
		if data, ok := <-channel; ok {
			fmt.Println(data)
		} else {
			break
		}
	} */

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("main end")
}
