package main

import "fmt"

func main() {
	fmt.Println("foo1 = ", foo1("foo1", 1))

	result1, result2 := foo2("foo2", 3)
	fmt.Println("result1 =", result1, "result2 =", result2)

	result1, result2 = foo3("foo3", 4)
	fmt.Println("result1 =", result1, "result2 =", result2)

	result1, result2 = foo4("foo4", 5)
	fmt.Println("result1 =", result1, "result2 =", result2)
}

func foo1(str string, num int) int {
	fmt.Println("str = ", str)
	fmt.Println("num = ", num)
	return num + 100
}

func foo2(str string, num int) (int, int) {
	fmt.Println("str = ", str)
	fmt.Println("num = ", num)
	return num * 2, num % 2
}

func foo3(str string, num int) (result1 int, result2 int) {
	fmt.Println("str = ", str)
	fmt.Println("num = ", num)

	result1 = num * 3
	result2 = result1 / 2

	return
}

func foo4(str string, num int) (result1, result2 int) {
	fmt.Println("str = ", str)
	fmt.Println("num = ", num)

	result1 = num * 4
	result2 = result1 / 2

	return
}
