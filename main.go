package main

import "fmt"

type Book struct {
	Name string
	Author string
}

func (book Book) String() string {
	return fmt.Sprintf("%s by %s", book.Name,book.Author)
}

func (book *Book) SetName(s string) {
	book.Name = s
}
 
func main() {
	book := Book{
		Name: "Harry Potter",
		Author: "J. K. Rowling",
	}

	fmt.Println(book.String()) // "Harry Potter by J. K. Rowling"
	book.SetName("new book")
	fmt.Printf("book: %v\n", book) // book: new book by J. K. Rowling
}
