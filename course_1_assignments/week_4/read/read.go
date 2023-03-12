package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Declare the string variables used to store user input.
	var filename string

	// Declare the Person struct
	type Person struct {
		fname string
		lname string
	}

	// Declare the slice
	sli := make([]Person, 0, 3)

	// Ask for user input.
	fmt.Println("Please enter the txt filename:")
	fmt.Scanln(&filename)

	// Open the text file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a scanner to read the file contents
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by space
		names := strings.Split(line, " ")

		// Extract first and lastname. Remove any spaces
		firstName := strings.ReplaceAll(names[0], " ", "")
		lastName := strings.ReplaceAll(names[1], " ", "")

		// Create a new Person struct
		newPerson := Person{
			fname: firstName,
			lname: lastName,
		}

		// Store the newPerson in the slice.
		sli = append(sli, newPerson)
	}

	// Close the file
	file.Close()

	// Print out the slice of Person structs
	fmt.Println(sli)

}
