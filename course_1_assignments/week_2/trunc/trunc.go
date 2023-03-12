package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Declare the string variables used to store user input.
	var userInput string

	// Loop twice, for 2 test cases.
	for i := 1; i <= 2; i++ {

		// Asks for a floating point number from the user.
		fmt.Printf("\nExecution number: %d\n", i)
		fmt.Println("Please enter a floating point number:")
		fmt.Scanln(&userInput)

		// Attempt to cast the input String to floating point string
		var floatNum float64
		floatNum, error := strconv.ParseFloat(userInput, 64)
		if error != nil {
			fmt.Println("Error: ", error)
			return
		}

		// Attempt to cast the float to a truncated int. Prints the truncated number.
		var intNum int64
		intNum = int64(floatNum)
		fmt.Printf("Truncated int is: %d\n", intNum)
	}
}
