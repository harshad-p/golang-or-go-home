package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (book Book) ToString() {
	fmt.Println("Title\t\t:\t", book.Title)
	fmt.Println("Author\t\t:\t", book.Author)
	fmt.Println("Pages\t\t:\t", book.Pages)
}

func (book Book) NotPossibleToResetPages() {
	book.Pages = 0
}

func (book *Book) ResetPages() {
	book.Pages = 0
}

func main() {
	harryPotterBook := Book{
		Title:  "Harry Potter and the Goblet of Fire",
		Author: "J. K. Rowling",
		Pages:  100,
	}

	// Print a struct directly
	fmt.Println(harryPotterBook)

	// Print a struct field
	fmt.Println(harryPotterBook.Pages)

	// Modify value of a field
	harryPotterBook.Pages = 700
	fmt.Println(harryPotterBook.Pages)

	// Use connected method to print Book
	harryPotterBook.ToString()

	// Try to modify pages using instance copy
	harryPotterBook.NotPossibleToResetPages()
	fmt.Println("After trying to modify pages using instance copy:")
	harryPotterBook.ToString()

	// Try to modify pages using pointer
	harryPotterBook.ResetPages()
	fmt.Println("After trying to modify pages using pointer:")
	harryPotterBook.ToString()
}
