package main

import (
	"fmt"
	"log"
)

func main() {
	//var whatToSay string
	//var saySomeThingElse string
	var i int

	//String Variables
	var sayMultiple string
	var world string

	whatToSay := saySomeThing("Hello Golang")
	log.Println(whatToSay)

	saySomeThingElse := saySomeThing("GoodBye")
	fmt.Println(saySomeThingElse)

	//if second parameter is unnecessary , we can remove it and use underscore mean (_).  
	sayMultiple, _ = saySomeThingMultiple("Hello");
	log.Println(sayMultiple);
	
	sayMultiple, world = saySomeThingMultiple("Golang")
	
	log.Println(world); // returns World
	log.Println(sayMultiple); // returns Golang
	log.Println(saySomeThingMultiple("Hello"))

	log.Println(saySomeThing("Finally"))

	//Integer Variables
	log.Println(i) // if dont assign any value returns 0 defaultly

	i = 44;
	log.Println(i);
}

func saySomeThing(s string) string {
	return s
}

func saySomeThingMultiple(s string) (string, string) {
	return s, "World"
}
