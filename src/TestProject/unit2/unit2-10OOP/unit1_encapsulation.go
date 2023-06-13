package main

import "fmt"

type Book struct {
	Title string
	Auth  string
}

func (book *Book) Show() {
	fmt.Printf("%+v\n", *book)
}

func (book *Book) GetAuth() string {
	return book.Auth
}

func (book *Book) SetAuth(auth string) {
	book.Auth = auth
}

func main() {
	book1 := Book{
		Title: "Golang",
		Auth:  "Boyen",
	}
	book1.Show()

	book1.SetAuth("Hsieh")
	fmt.Println(book1.GetAuth())
}
