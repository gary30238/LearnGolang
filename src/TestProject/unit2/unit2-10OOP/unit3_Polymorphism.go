package main

import "fmt"

type IAnimal interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat sleep")
}

func (cat *Cat) GetColor() string {
	return cat.color
}

func (cat *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (dog *Dog) Sleep() {
	fmt.Println("Dog sleep")
}

func (dog *Dog) GetColor() string {
	return dog.color
}

func (dog *Dog) GetType() string {
	return "Dog"
}

func ShowInfo(animal IAnimal) {
	animal.Sleep()
	// fmt.Println("color =", animal.GetColor(), "type =", animal.GetType())
}

func main() {
	var animal IAnimal
	cat := Cat{"White"}
	animal = &cat
	animal.Sleep()

	dog := Dog{"Yellow"}
	animal = &dog
	animal.Sleep()

	ShowInfo(&cat)
	ShowInfo(&dog)
}
