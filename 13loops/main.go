package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// i here will be an index
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("value is %v\n", day)
	// }

	rogueValue := 1

	// for rogueValue < 10 {

	// 	if rogueValue == 5 {
	// 		break
	// 	}

	// 	fmt.Printf("Value is %v\n", rogueValue)
	// 	rogueValue++
	// }

	for rogueValue < 10 {

		if rogueValue == 2 {
			// provide the label name after goto
			goto lco
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Printf("Value is %v\n", rogueValue)
		rogueValue++
	}

	// label - use nay name for label except reserved keywords
lco:
	fmt.Println("Jumped to learn code statement")

}
