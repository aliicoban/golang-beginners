package main

import (
	"log"
	"time"
)
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {

	user := User{
		FirstName : "Ali",
		LastName :  "Coban",
		PhoneNumber : "555 555 55 55",
		Age : 26,
	}
	log.Println(user);
}
