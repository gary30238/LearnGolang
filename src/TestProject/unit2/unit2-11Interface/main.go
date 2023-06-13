package main

import "fmt"

// interface可儲存任意型別的數值
func checkType(arg interface{}) {
	switch value := arg.(type) {
	case string:
		fmt.Printf("Type of the %v = String\n", value)
	case int:
		fmt.Printf("Type of the %v = Int\n", value)
	case Book:
		fmt.Printf("Type of the %v = Book\n", value)
	default:
		fmt.Printf("Type of the %v = Other type\n", value)
	}
}

type Book struct {
	Name string
	Auth string
}

func main() {
	book := Book{"Golang", "Boyen"}

	checkType(book)
	checkType(100)
	checkType("test")
	checkType(3.14)
}
