package main

import (
	"fmt"
	"sync"
)

/*
In this file, there are 2 goroutines being created, A and B. Both of these goroutines will increment a global counter variable 5 times. A is scheduled before B.

However, when executing, B is executed before A. Also, the counter values overlap with each other. One such output is as follows:
Name:  Goroutine B . Counter:  1
Name:  Goroutine B . Counter:  3
Name:  Goroutine B . Counter:  4
Name:  Goroutine B . Counter:  5
Name:  Goroutine B . Counter:  6
Name:  Goroutine A . Counter:  2
Name:  Goroutine A . Counter:  7
Name:  Goroutine A . Counter:  8
Name:  Goroutine A . Counter:  9
Name:  Goroutine A . Counter:  10
Main is done.

The Race condition occurs as both Goroutines A and B are trying to access the same variable, `global_counter`.

*/

var global_counter = 0

// Accesses the global variable and keeps increments it twice
func increment(name string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		temp := global_counter
		temp++
		global_counter = temp
		fmt.Println("Name: ", name, ". Counter: ", global_counter)
	}
	wg.Done()
}

func main() {

	// Create a sync.WaitGroup
	var wg sync.WaitGroup

	// Creates the 2 goroutines
	wg.Add(2)
	go increment("Goroutine A", &wg)
	go increment("Goroutine B", &wg)

	// Waits until the 2 goroutines have completed.
	wg.Wait()
	fmt.Println("Main is done.")

}


