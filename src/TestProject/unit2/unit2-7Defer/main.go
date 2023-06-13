package main

import "fmt"

func deferFunc() {
	fmt.Println("deferFunc Called...")
}

func returnFunc() int {
	fmt.Println("returnFunc Called...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// 依序放入stack內,等func執行完(生命週期)再取出來
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")

	fmt.Println("main called...")
	// defer會在return之後執行,等同於func執行完(生命週期)再取出來
	returnAndDefer()
}
