package main

import (
	"github.com/alicobanserver/helpers"
	"log"
)

func main() {
	log.Println("Hello")
	var myVar helpers.SomeType

	myVar.TypeName = "Some name"
	myVar.TypeNumber = "Some number"

	log.Println(myVar)

}
