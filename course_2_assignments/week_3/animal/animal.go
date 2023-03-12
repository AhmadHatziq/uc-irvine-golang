package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {

	// Declare variables
	var userInput string
	cowAnimal := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	birdAnimal := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snakeAnimal := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	// Add aninals to a map
	animals := make(map[string]Animal)
	animals["cow"] = cowAnimal
	animals["bird"] = birdAnimal
	animals["snake"] = snakeAnimal

	// Get input from user in a continuous loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		if scanner.Scan() {

			// Read in line
			userInput = scanner.Text()

			// Parse the 2 words
			animal_entered := strings.Split(userInput, " ")[0]
			action_entered := strings.Split(userInput, " ")[1]

			// Extract the matching animal
			matchingAnimal, ok := animals[animal_entered]
			if !ok {
				fmt.Println("Animal not found. ")
				continue
			}

			// Extract the method based on the action entered by the user
			switch action_entered {
			case "eat":
				matchingAnimal.Eat()
			case "speak":
				matchingAnimal.Speak()
			case "move":
				matchingAnimal.Move()
			default:
				fmt.Println("Invalid move entered")
				continue
			}

		}
		// fmt.Println(userInput)
	}

}
