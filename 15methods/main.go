package main

import "fmt"

func main() {

	akshay := User{"Akshay", "ak@gmail.com", true, 16}
	fmt.Println(akshay)

	akshay.GetStatus()
	akshay.NewMail()
	fmt.Println(akshay)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
