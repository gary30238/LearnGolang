package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string
	Age    int
	Gender bool
}

type Class struct {
	Id       string
	Students []Student
}

func main() {
	s := Student{"Boyen", 18, true}
	c := Class{"1班", []Student{s, s, s}}

	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Println(string(bytes))

	var c2 Class
	err = json.Unmarshal(bytes, &c2)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("%+v\n", c2)
}
