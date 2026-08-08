package main

import "fmt"

func main() {
	type Book struct {
		Title  string
		Author string
		Pages  int
	}

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
}
