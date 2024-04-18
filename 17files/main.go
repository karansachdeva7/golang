package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - Learn to code in go"

	file, err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	panic(err)
	// }00
	checkNilErr(err)

	// byte writer
	// copy buffer
	// good utility methods

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./mygofile.txt")
}

func readFile(filename string) {
	// creation is os operation
	// but for reading/manipulating the file, there's diff utility given to us, ioutil

	// data is read in byte format
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		// shutdown execution of program and will show what error happended
		panic(err)
	}
}
