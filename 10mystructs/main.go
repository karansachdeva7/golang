package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	karan := User{"karan", "karansample@go.dev", true, 20}
	fmt.Println(karan)

	fmt.Printf("karan details are: %+v\n", karan)
	fmt.Printf("Name is %v and Email is %v", karan.Name, karan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
