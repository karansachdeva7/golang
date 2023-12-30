package main

import "fmt"

var h = "hello"

// Public
const LoginToken string = "some value"

func main() {
	var username string = "Karan"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type %T \n", smallValue)

	var smallFloat float32 = 255.5577867878644554455
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	var bigFloat float64 = 255.5577867878644554455
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type %T \n", bigFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T \n", anotherVariable)

	// implicit type/way of declaring variables
	var website = "www.google.com"
	fmt.Println(website)

	// no var style
	numberOfUsers := 10000
	fmt.Println(numberOfUsers)

	// Note: Inside the method, you are allowed to use walrus operator
	// but not allowed globally (outside function)
	// you can use syntax- var user = "Karan"

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)
}
