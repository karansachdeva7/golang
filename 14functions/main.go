package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	// not allowed to write func inside func
	// func greeterTwo() {
	// 	fmt.Println("Another method")
	// }
	// greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proRes, _ := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proRes)
}

func adder(a, b int) int {
	return a + b
}

func greeter() {
	fmt.Println("Hello from Golang")
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "result processed"
}
