package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 屬性可用“加入tag註釋
type Human struct {
	Name string `info:"name" doc:"名字"`
	Sex  string `info:"sex" doc:"性別"`
}

// jsonstring默認使用變量名 可用json替換變量名
type Movie struct {
	Title  string   `json:"movieName"`
	Year   int      `json:"movieYear"`
	Price  int      `json:"NT"`
	Actors []string `json:"movieActors"`
}

func reflectTag(arg interface{}) {
	argType := reflect.TypeOf(arg)

	for i := 0; i < argType.NumField(); i++ {
		tagInfo := argType.Field(i).Tag.Get("info")
		tagDoc := argType.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagInfo, "doc:", tagDoc)
	}
}

func main() {
	var human Human

	reflectTag(human)

	movie := Movie{"玩命關頭特別行動", 2023, 300, []string{"胖光頭", "瘦光頭"}}

	// 序列化
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Println(string(jsonStr))

	// 反序列化
	var myMovie Movie
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json umMarshal error", err)
		return
	}

	fmt.Println(myMovie.Title)
	fmt.Println(myMovie.Year)
	fmt.Println(myMovie.Price)
	fmt.Println(myMovie.Actors)
}
