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

func (a *Animal) Eat() {
	fmt.Println(a.food)
}
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}
func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	var req string
	for {
		fmt.Printf("\n%v", ">")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			req = scanner.Text()

		}
		var lis []string
		lis = strings.Split(req, " ")
		if lis[0] == "exit" {
			break
		}

		switch lis[0] {
		case "cow":
			switch lis[1] {
			case "speak":
				cow.Speak()
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			}
		case "bird":
			switch lis[1] {
			case "speak":
				bird.Speak()
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			}
		case "snake":
			switch lis[1] {
			case "speak":
				snake.Speak()
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			}
		default:
			fmt.Println("Wrong input")
		}
	}
}
