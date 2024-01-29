package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in goalng")

	// mention size explicitly (compulsion)
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	// depends on reserved mem size, i.e. output 4 below, not actual values
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is:", vegList)
	fmt.Println("Vegy list is:", len(vegList))

}
