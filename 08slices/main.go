package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	// fmt.Println(fruitList)

	// memory mgmt - two ways to allocate memory

	// default allocation of mem
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777

	// the moment you use append, mem reallocated, and all values will be accomodated
	highScores = append(highScores, 555, 666)

	// fmt.Println(highScores)

	sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	// use append also to remove the value
	// reallocates the mem
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
