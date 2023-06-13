package main

import (
	"fmt"
)

// sub goroutine
/* func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new Goroutine:", i)
		time.Sleep(1 * time.Second)
	}
} */

// main goroutine
func main() {
	/* go newTask()

	i := 0
	for {
		i++
		fmt.Println("main goroutine:", i)
		time.Sleep(1 * time.Second)
	} */

	/* go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			// 只會退出當前func
			return
			// 退出當前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}() */

	// 無法取得返回值
	go func(a int, b int) bool {
		fmt.Println(a + b)
		return true
	}(10, 20)

	fmt.Scanln()
}
