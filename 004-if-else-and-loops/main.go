package main

import "fmt"

func adultOrMinor(age int) string {
	if age >= 18 {
		return "Adult"
	} else {
		return "Minor"
	}
}

func main() {
	var age1 int = 15
	var age2 int = 21
	fmt.Println("Age =", age1, "=>", adultOrMinor(age1))
	fmt.Println("Age =", age2, "=>", adultOrMinor(age2))
}
