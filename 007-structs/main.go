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
		Pages:  1000,
	}

	// Print a struct directly
	fmt.Println(harryPotterBook)
}
