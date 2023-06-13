package main

import "fmt"

// global variables
// 1
// var unitName = "Unit2-1Variable"
// 2
var (
	unitNo   = "Unit2-1"
	unitName = "Variable"
)

func main() {
	fmt.Println(unitNo + ":" + unitName)
	// local variables
	// 變數聲明/賦值/使用
	// 1
	var num1 int
	// 不賦值 default value = 0
	num1 = 1
	fmt.Println("num1:", num1)

	// 2
	var num2 int = 2
	fmt.Println("num2:", num2)

	// 3
	var num3 = 3
	fmt.Println("num3:", num3)

	// 4 only local variable
	num4 := 4
	fmt.Println("num4:", num4)

	fmt.Println("---------------------------------------------")

	// 多變數聲明
	var num5, num6 int
	fmt.Println(num5)
	fmt.Println(num6)

	var name, heigth, weight = "boyen", 175, 70
	fmt.Println(name)
	fmt.Println(heigth)
	fmt.Println(weight)

	name1, heigth1, weight1 := "ying", 160, 45
	fmt.Println(name1)
	fmt.Println(heigth1)
	fmt.Println(weight1)

	fmt.Println("---------------------------------------------")
	// data type : bool / numeric / string / byte
	var (
		// default = false
		bool1 bool
		// int8 / int16 / int32 / int64
		// unit8 / unit16 / unit32 / unit64
		// default = 0
		int1 int
		// float32 / float64
		float1 float32
		float2 float32 = 3.141
		// default = ""
		string1 string
		char1   = 'a'
	)
	/*
		default format:%v
		type:%T
		bool:%t
		int:%d %nd 10進制
			%b 2進制
			%o 8進制
			%x 16進制
			%X 16進制(大寫)
			%U unicode
		float:%g
			  %.nf
			  %w.nf
		string:%s %ns
		char:%c
	*/
	fmt.Printf("bool: %v\nint: %v\nfloat : %v\nstring: %v\n", bool1, int1, float1, string1)
	fmt.Printf("char: %c\n", char1)
	fmt.Printf("type of bool1: %T\n", bool1)
	fmt.Printf("float format: %.2f", float2)
	fmt.Printf("float format: %5.2f", float2)
}
