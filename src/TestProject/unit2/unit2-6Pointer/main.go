package main

import "fmt"

func swap(pa *int, pb *int) {
	var temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 10
	var b int = 20

	// a, b = b, a
	swap(&a, &b)

	fmt.Println("a =", a, "b =", b)

	// 一級指針 *type
	var p = &a

	fmt.Println("pointer of a =", &a)
	fmt.Println("value of p =", p)

	// 二級指針 **type
	var pp = &p

	fmt.Println("pointer of p =", &p)
	fmt.Println("value of pp =", pp)
}
