package main

import (
	"fmt"
	"strings"
)

func main() {

	// Declare the string variables used to store user input.
	var userInput string

	// Ask for user input.
	fmt.Println("Please enter the string:")
	fmt.Scanln(&userInput)

	// Searches for 'i' and 'a' and 'n' in the string.
	isAinside, startWithI, endWithN := false, false, false
	isAinside = strings.Contains(userInput, "a")
	startWithI = strings.HasPrefix(userInput, "i")
	endWithN = strings.HasSuffix(userInput, "n")

	if (isAinside) && (startWithI) && (endWithN) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
