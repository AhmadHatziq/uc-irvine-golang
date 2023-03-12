package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Goroutine function to sort an array.
func sort_array(name string, array_to_sort []int, wg *sync.WaitGroup) {

	// Print the elements before sorting.
	fmt.Println("Goroutine: ", name, ".Unsorted elements: ", array_to_sort)

	// Sort the array
	sort.Sort(sort.IntSlice(array_to_sort))

	// Print the elements after sorting
	fmt.Println("Goroutine: ", name, ".Sorted elements: ", array_to_sort)

	wg.Done()

}

func main() {

	// Get input of numbers from the user.
	var userInput string
	fmt.Println("Please enter the numbers seperated by a space: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()

	}
	number_strings := strings.Split(userInput, " ")
	numbers := make([]int, len(number_strings))
	for i, number_string := range number_strings {
		number, err := strconv.Atoi(number_string)
		if err != nil {
			panic(err)
		}
		numbers[i] = number
	}

	// Partition the numbers into 4 subarrays
	size := len(numbers) / 4
	remainder := len(numbers) % 4
	start := 0
	partitions := make([][]int, 4) // Make a slice which contains 4 elements, where each is another slice on integers
	for i := 0; i < 4; i++ {
		end := start + size
		if i < remainder {
			end++
		}
		partitions[i] = numbers[start:end]
		start = end
	}

	// Create a sync.WaitGroup
	var wg sync.WaitGroup

	// Call 4 subroutines to sort the array
	wg.Add(4)
	go sort_array("A", partitions[0], &wg)
	go sort_array("B", partitions[1], &wg)
	go sort_array("C", partitions[2], &wg)
	go sort_array("D", partitions[3], &wg)

	// Waits until the 4 goroutines have completed and sort the final array.
	wg.Wait()
	final_sorted_array := merge(partitions[0], partitions[1], partitions[2], partitions[3])
	fmt.Println("Final sorted array: ", final_sorted_array)

}

// Used to sort and merge arrays
func merge(arrays ...[]int) []int {
	var result []int

	for _, arr := range arrays {
		result = append(result, arr...)
	}

	sort.Sort(sort.IntSlice(result))

	return result
}
