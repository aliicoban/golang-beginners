//Loops and ranging over data
package main

import (
	"log"
)

type User struct {
	FirstName string
}

func main() {
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	//range
	mySlice := []string{"dog", "cat", "horse", "fish", "banana", "apple"}

	for index, value := range mySlice {
		log.Println(index, value)
	}

	// if we dont use index
	for _, value := range mySlice {
		log.Println(value)
	}

	//usage with map

	myMap := make(map[string]string)
	myMap["dog"] = "Dog"
	myMap["cat"] = "Cat"
	myMap["fish"] = "Fish"

	for _, value := range myMap {
		log.Println(value)
	}

	//usage with struct

	var myUser []User
	u1 := User{
		FirstName: "Ali",
	}
	u2 := User{
		FirstName: "Huawei",
	}
	myUser = append(myUser, u1)
	myUser = append(myUser, u2)

	for _, value := range myUser {
		log.Println(value)
		log.Println(value.FirstName)
	}

}
