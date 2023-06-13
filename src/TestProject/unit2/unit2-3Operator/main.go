package main

import "fmt"

func main() {
	// 加減
	fmt.Println(+10)
	fmt.Println(4 + 7)
	fmt.Println(5 - 9)
	fmt.Println("abc" + "def")

	// 乘 / 除 / 餘
	fmt.Println(5 * 3)
	fmt.Println(10 / 3)
	fmt.Println(10.0 / 3)
	fmt.Println(10 % 3)

	// Assignment operators
	n1 := 10
	n1++
	fmt.Println(n1)
	n2 := 20
	n1 += n2
	fmt.Println(n1)
	// change
	a := 4
	b := 8
	a, b = b, a
	fmt.Println(a, b)

	// &取指標位址 *取指標數值
	age := 18
	fmt.Println(&age)
	prt := &age
	fmt.Println(&prt)
	fmt.Println(*prt)
}
