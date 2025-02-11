package main

import (
	"fmt"
)

func main() {
	// Ask the user to enter an integer
	fmt.Println("Enter an integer to check if it's even or odd:")

	var number int
	// Read the input number
	fmt.Scan(&number)

	// Check if the number is even or odd
	if number%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}
