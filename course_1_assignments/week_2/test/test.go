package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (this *Cow) Eat() {
	fmt.Println("grass")
}

func (this *Cow) Move() {
	fmt.Println("walk")
}

func (this *Cow) Speak() {
	fmt.Println("moo")
}

func (this *Bird) Eat() {
	fmt.Println("worms")
}

func (this *Bird) Move() {
	fmt.Println("fly")
}

func (this *Bird) Speak() {
	fmt.Println("peep")
}

func (this *Snake) Eat() {
	fmt.Println("mice")
}

func (this *Snake) Speak() {
	fmt.Println("hsss")
}

func (this *Snake) Move() {
	fmt.Println("slither")
}

type animals_t map[string]Animal

func handleQuery(name string, action string, animals animals_t) {
	a := animals[name]
	switch action {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func handleNewAnimal(name string, specie string, animals animals_t) {
	var a Animal
	switch specie {
	case "cow":
		c := Cow{}
		a = &c
	case "snake":
		s := Snake{}
		a = &s
	case "bird":
		b := Bird{}
		a = &b
	}
	animals[name] = a
}

func main() {
	var arg1, arg2, arg3 string
	animals := make(animals_t, 0)
	fmt.Println("ctrl c to interrupt")
	for true {
		fmt.Print("> ")
		fmt.Scan(&arg1, &arg2, &arg3)
		if arg1 == "newanimal" {
			handleNewAnimal(arg2, arg3, animals)
		} else if arg1 == "query" {
			handleQuery(arg2, arg3, animals)
		} else {
			fmt.Println("Unrecognized action")
		}
	}

}
