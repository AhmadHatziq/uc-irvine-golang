package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define Animal interface.
type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

// Define Cow type and methods which satisfies the Animal interface
type Cow struct {
	name string
}

func (c Cow) GetName() string {
	return c.name
}
func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Define Bird type and methods which satisfies the Animal interface.
type Bird struct {
	name string
}

func (b Bird) GetName() string {
	return b.name
}
func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Define Snake type and methods which satisfies the Animal interface.
type Snake struct {
	name string
}

func (s Snake) GetName() string {
	return s.name
}
func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	// Declare variables
	var userInput string
	animalSlice := make([]Animal, 0)

	// Get input from user in a continuous loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		if scanner.Scan() {

			// Read in line
			userInput = scanner.Text()
			userInput = strings.ToLower(userInput)

			// Parse the 3 words
			enteredWords := strings.Split(userInput, " ")
			if len(enteredWords) != 3 {
				fmt.Println("Invalid command")
				continue
			}

			// There are only 2 possible actions, based on the first string
			firstString := enteredWords[0]
			switch firstString {
			case "newanimal":

				// Extract name and animal type
				animalName := enteredWords[1]
				animalType := enteredWords[2]

				// Create the animal with that name, using the Animal interface.
				// Stores the animal into the animal slice
				switch animalType {
				case "cow":
					animal := Cow{name: animalName}
					animalSlice = append(animalSlice, animal)
				case "bird":
					animal := Bird{name: animalName}
					animalSlice = append(animalSlice, animal)
				case "snake":
					animal := Snake{name: animalName}
					animalSlice = append(animalSlice, animal)
				default:
					fmt.Println("Invalid command")
					continue
				}

				// Print back to the user
				fmt.Println("Created it!")

			case "query":

				// Get the animal name and action
				animalName := enteredWords[1]
				animalAction := enteredWords[2]

				// Iterate and find matching animal
				var matchingAnimal Animal
				for _, a := range animalSlice {
					if a.GetName() == animalName {
						matchingAnimal = a
					}
				}

				// Return error if no matching animal is found
				if matchingAnimal == nil {
					fmt.Println("Animal name not found")
					continue
				}

				// Call the method based on the action specified.
				switch animalAction {
				case "eat":
					matchingAnimal.Eat()
				case "speak":
					matchingAnimal.Speak()
				case "move":
					matchingAnimal.Move()
				default:
					fmt.Println("Action not found")
				}

			default:
				fmt.Println("Invalid command entered")
				continue
			}

		}
		// fmt.Println(userInput)
	}

}
