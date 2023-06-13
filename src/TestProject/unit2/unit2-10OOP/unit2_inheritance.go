package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func (people *People) Talk(saying string) {
	fmt.Println(people.Name, "say", saying)
}

func (people *People) Eat(food string) {
	fmt.Println(people.Name, "eat", food)
}

type Superman struct {
	// 繼承自People
	People

	Level int
}

func (superman *Superman) Talk() {
	fmt.Println(superman.Name, "say I'm Superman")
}

func (superman *Superman) Fly() {
	fmt.Println(superman.Name, "fly...")
}

func main() {
	people1 := People{
		"Boyen",
		30,
	}
	people1.Talk("Hi")
	people1.Eat("Something")

	/* superman1 := Superman{
		People{
			"Clark",
			30,
		},
		100,
	} */
	var superman1 Superman
	superman1.Name = "Clark"
	superman1.Age = 30
	superman1.Level = 100
	superman1.Talk()
	superman1.Eat("Monster")
	superman1.Fly()
}
