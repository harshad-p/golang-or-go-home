package main

import "fmt"

func divide(dividend, divisor float32) (float32, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("Cannot divide by 0.")
	}

	return dividend / divisor, nil
}

func printDivisionResult(dividend, divisor, result float32, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%f / %f = %f\n", dividend, divisor, result)
	}
}

func main() {
	result1, error1 := divide(4, 2)
	printDivisionResult(4, 2, result1, error1)

	result2, error2 := divide(4, 0)
	printDivisionResult(4, 2, result2, error2)
}
