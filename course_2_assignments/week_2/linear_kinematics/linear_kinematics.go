package main

import (
	"fmt"
	"strconv"
)

// This function is used to ask the 'term' to the user and returns back the float.
func process_input(term string) float64 {

	// Prompt the user to enter the term
	fmt.Printf("Please enter the value for %s: ", term)

	// Process the input
	var userInput string
	fmt.Scanln(&userInput)
	float_value, _ := strconv.ParseFloat(userInput, 64)

	// Return value
	return float_value
}

// Note that the 2nd 'func' must be on the same line.
// This is a test function not used for the assessment.
func add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 0.50*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
	return fn
}

func main() {

	// Ask for acceleration, initial velocity, initial displacement & time from the user.
	acceleration := process_input("acceleration")
	initialVelocity := process_input("initial velocity")
	initialDisplacement := process_input("initial displacement")
	time := process_input(("time"))

	fmt.Println(acceleration, initialVelocity, initialDisplacement, time)

	// Generate the displacement function of time.
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	// Generate and print the displacement after the input time.
	fmt.Printf("Displacement is : %.2f \n", fn(time))

}
