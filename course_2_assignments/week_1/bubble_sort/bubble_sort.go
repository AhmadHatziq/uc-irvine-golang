package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(sli []int, i int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}

func BubbleSort(sli []int) {
	n := len(sli)

	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {

			if sli[j] > sli[j+1] {

				Swap(sli, j)

			}
		}
	}
}

func main() {

	// Declare variables
	var userInput string

	// Get input of integers from the user
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input the integers, seperated by space (eg 1 2 3 4 5):")
	if scanner.Scan() {
		userInput = scanner.Text()
	}

	// Store the list of integers into a slice
	integer_slice := make([]int, 0)
	numbers := strings.Split(userInput, " ")
	for _, element := range numbers {
		number, _ := strconv.Atoi(element)
		integer_slice = append(integer_slice, number)
	}

	// Sort the integer slice
	BubbleSort(integer_slice)

	// Prints the sorted slice
	fmt.Println(integer_slice)

}
