package main

import "fmt"

// Sometimes when you pass on variables, a copy of these
// variables is being passed on
// Whenever there is a case when you want to avoid such things to happen
// and you want absolute guarantee that always the actual value
// should be passed on, then we prefer that a pointer should be passed on,
// A pointer is a direct reference to the memory address,
// Since you are directly passing memory address, it makes hundred percent guarantee
// that actual value is being passed on.
func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber // a reference means ampersand
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)

}
