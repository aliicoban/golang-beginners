package main

import (
	"log"
)

func main() {
	var myString string
	myString = "Green"

	log.Println("myString set to:", myString)

	log.Println("Location is:", &myString) // pass a reference to this variable, to this function.
	changeUsingPointer(&myString)

	log.Println("After func call myString set to:", myString);
}

func changeUsingPointer(s *string) {
	// asteriks for actual pointer itself when you are pointing to location in memory.
	newValue := "Red"
	*s = newValue
}
