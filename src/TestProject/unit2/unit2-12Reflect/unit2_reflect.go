package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%w\n", user)
}

func GetTypeAndMethod(arg interface{}) {
	// type
	argType := reflect.TypeOf(arg)
	fmt.Println("ArgType is :", argType.Name())

	// value
	argValue := reflect.ValueOf(arg)
	fmt.Println("ArgValue is :", argValue)

	// field
	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)
		value := argValue.Field(i)

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// method
	for i := 0; i < argType.NumMethod(); i++ {
		m := argType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func main() {
	user := User{1, "Boyen", 30}

	GetTypeAndMethod(user)
}
