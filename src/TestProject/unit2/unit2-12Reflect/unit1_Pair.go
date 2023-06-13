package main

import (
	"fmt"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (book *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (book *Book) WriteBook() {
	fmt.Println("Write a book")
}

func main() {
	// pair<type:string value:"">
	var s string
	// pair<type:string value:"today">
	s = "today"

	// pair<type: value:>
	var allType interface{}
	// pair<type: string value:"today">
	allType = s

	str, ok := allType.(string)
	fmt.Println(str, ok)

	// pair<type:Book value:&Book{}>
	book := &Book{}

	var r Reader
	// pair<type:Book value:&Book{}>
	r = book
	r.ReadBook()

	var w Writer
	// pair<type:Book value:&Book{}>
	w = r.(Writer)
	w.WriteBook()
}
