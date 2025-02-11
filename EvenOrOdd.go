package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Print("\nEnter an integer to check if it's even or odd (or type 'exit' to quit): ")

		var input string
		fmt.Scan(&input)

		// Exit condition
		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		// Convert input to integer
		var number int
		_, err := fmt.Sscanf(input, "%d", &number)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		// Check even or odd
		if number%2 == 0 {
			fmt.Println("The number is even.")
		} else {
			fmt.Println("The number is odd.")
		}
	}
}
