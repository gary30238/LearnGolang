package main

import "fmt"

/* golang默認都是值傳遞
只有slice, map, channel
是用引入傳遞(pointer) */

// 傳入參數同樣要宣告長度
func printArr(arr [4]int) {
	for index, value := range arr {
		fmt.Println("index =", index, "value =", value)
	}
	// 值傳遞,值不會變化
	arr[0] = 100
}

// 傳入參數同樣要宣告長度
func printSlice(slice []int) {
	// 可用 _ 替代不須使用的變數
	for _, value := range slice {
		fmt.Println("value =", value)
	}
	// 引用傳遞,值會變化
	slice[0] = 100
}

func main() {
	// 固定長度的陣列array
	fmt.Println("Array")
	var arr1 [10]int
	arr2 := [10]int{1, 2, 3, 4}
	arr3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for index, value := range arr2 {
		fmt.Println("index =", index, "value =", value)
	}

	printArr(arr3)
	// 查看在func內改變的值,是否有變化
	for index, value := range arr3 {
		fmt.Println("index =", index, "value =", value)
	}

	fmt.Printf("type of arr1 =%T\n", arr1)
	fmt.Printf("type of arr2 =%T\n", arr2)
	fmt.Printf("type of arr3 =%T\n", arr3)

	fmt.Println("========================================")

	// 動態陣列slice
	fmt.Println("Slice")
	/* var slice1 []int
	slice1 := make([]int, 4) */
	slice1 := []int{1, 2, 3, 4}
	printSlice(slice1)
	// 查看在func內改變的值,是否有變化
	for index, value := range slice1 {
		fmt.Println("index =", index, "value =", value)
	}

	// len = 3, cap = 4
	slice2 := make([]int, 3, 4)
	fmt.Printf("len = %v, cap = %v, slice = %v\n", len(slice2), cap(slice2), slice2)

	// len = 4, cap = 4, [0, 0, 0, 1]
	slice2 = append(slice2, 1)
	fmt.Printf("len = %v, cap = %v, slice = %v\n", len(slice2), cap(slice2), slice2)

	// 當length > cap, 會依slice當前cap追加相同cap
	// len = 5, cap = 8, [0, 0, 0, 1, 2]
	slice2 = append(slice2, 2)
	fmt.Printf("len = %v, cap = %v, slice = %v\n", len(slice2), cap(slice2), slice2)

	var slice3 []int
	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2)
	slice3 = append(slice3, 3)
	fmt.Printf("len = %v, cap = %v, slice = %v\n", len(slice3), cap(slice3), slice3)

	// 擷取slice
	slice4 := slice3[0:2]
	fmt.Println(slice4)

	// 複製
	var slice5 []int
	copy(slice5, slice3)
	fmt.Println(slice5)
}
