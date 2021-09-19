package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMapStr := make(map[string]string)
	myMapInt := make(map[string]int)
	myMapInt2 := make(map[int]int)
	myMapFloat := make(map[string]float32)

	myMapStr["cat"] = "Cookie"
	log.Println(myMapStr["cat"])

	myMapInt["age"] = 26
	log.Println(myMapInt["age"])

	myMapInt2[1] = 2
	log.Println(myMapInt2[1])

	myMapFloat["number"] = 11.1
	log.Println(myMapFloat["number"])

	//maps with struct

	myMapStruct := make(map[string]User)

	me := User{
		FirstName: "Ali",
		LastName:  "Coban",
	}
	myMapStruct["me"] = me
	log.Println(myMapStruct["me"].FirstName, myMapStruct["me"].LastName)
}
