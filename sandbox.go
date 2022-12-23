package main

import "fmt"

func main() {
	//var name string = "I am a string"
	name := "I am a string" // these two lines are 100% equivalent.
	name = "I am replacing the last string"
	fmt.Println(name)

	// var isActive bool = false
	isActive := false
	fmt.Println(isActive)

	// var age int = 36
	age := 36
	fmt.Println(age)

	// var price float64 = 5.43
	price := 5.43
	fmt.Println(price)

	const left = "LEFT" // cannot declare using := and cannot be reassigned
	fmt.Println(left)

	//Arrays
	movies := []string{"Home Alone", "Die Hard", "John Wick", "Mask"}

	// go has only one loop (for) construct.
	for i := 0; i < len(movies); i++ {
		fmt.Println(movies[i])
	}

	// Another way to loop using a range
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, number := range numbers {
		fmt.Println(number)
	}

	// Forever for loop
	for {
		fmt.Println("Hello!")
	}
}
