package main

import "fmt"

const (
	// iota初始=0,依序累加,只可與const搭配使用
	Taipei = iota
	Taichung
	Tainan
)
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
	g, h = iota * 2, iota * 3
	i, j
)

func main() {
	const unitName = "unit2-2Const"
	fmt.Println(unitName)

	fmt.Println("Taipei: ", Taipei)
	fmt.Println("Taichung: ", Taichung)
	fmt.Println("Tainan: ", Tainan)

	fmt.Println("a: ", a, "b: ", b)
	fmt.Println("c: ", c, "d: ", d)
	fmt.Println("e: ", e, "f: ", f)
	fmt.Println("g: ", g, "h: ", h)
	fmt.Println("i: ", i, "j: ", j)
}
