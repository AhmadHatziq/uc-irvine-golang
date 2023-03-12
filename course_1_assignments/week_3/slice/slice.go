package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// Declare the variables used to store user input.
	var userInput string

	// Declare the empty integer slice of size (length) 3
	sli := make([]int, 0, 3)

	// Create an infinite loop that will only exit if the user enters a 'x'.
	for {

		var userInt int

		// Retrieve input from user
		fmt.Println("Please enter an int or enter X to exit:")
		fmt.Scanln(&userInput)

		// Exits out if user inputs a 'x'
		if strings.Contains(strings.ToLower(userInput), "x") {
			fmt.Println("Exiting loop.")
			break
		}

		// Casts the string input to a integer.
		userInt, error := strconv.Atoi(userInput)
		if error != nil {
			fmt.Println("Error: ", error)
			break
		}

		// Store int to the slice
		sli = append(sli, userInt)

		// sorts the slice/array
		sort.Ints(sli)

		// Print out slice contents
		fmt.Println(sli)

	}

}
