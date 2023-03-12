package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// Declare the variables used to store user input.
	var userInput string

	// Read in user name
	fmt.Println("Please enter your name:")
	fmt.Scanln(&userInput)
	userName := userInput

	// Read in address
	fmt.Println("Please enter your address: ")
	fmt.Scanln(&userInput)
	userAddress := userInput

	// Create a new map with the keys: name, address
	userAddressMap := make(map[string]string)

	// Insert the 2 entries to the map
	userAddressMap["name"] = userName
	userAddressMap["address"] = userAddress

	// Marshal a JSON object from the map
	byteArray, error := json.Marshal(userAddressMap)
	if error != nil {
		fmt.Println(error)
		return
	}

	// Print back the JSON
	fmt.Println(byteArray)
	fmt.Println(string(byteArray))

}
