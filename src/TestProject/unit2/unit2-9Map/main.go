package main

import "fmt"

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key =", key, "value =", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	var map1 = make(map[int]string, 10)

	map1[1] = "C#"
	map1[2] = "python"
	map1[3] = "go"

	fmt.Println(map1)

	map2 := make(map[int]string)

	map2[1] = "html"
	map2[2] = "css"
	map2[3] = "js"

	fmt.Println(map2)

	map3 := map[int]string{
		1: "angular",
		3: "react",
		2: "vue",
	}

	fmt.Println(map3)

	cityMap := map[string]string{
		"China": "Beijing",
		"Japan": "Tokyo",
		"USA":   "NewYork",
	}

	printMap(cityMap)

	cityMap["China"] = "World"

	delete(cityMap, "China")

	changeValue(cityMap)
}
