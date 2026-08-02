package main

import "fmt"

func sayHello() {
	fmt.Println("Hello!")
}

func add(a int, b int) int {
	return a + b
}

func main() {
	a := 2
	b := 3

	sayHello()
	fmt.Println("Operations on", a, "and", b)
	fmt.Println("Addition:", add(a, b))
}
