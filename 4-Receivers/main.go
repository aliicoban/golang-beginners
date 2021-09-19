//Structs with Functions
package main

import (
	"log"
)

type myStruct struct {
	FirstName string
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Ali"

	myVar2 := myStruct{
		FirstName: "Huawei",
	}

	log.Println("myVar is :", myVar.FirstName);
	log.Println("myVar2 is :", myVar2.FirstName);

	log.Println("-------------");

	// with receiver function:
	log.Println("myVar is :", myVar.printFirstName());
	log.Println("myVar2 is :", myVar2.printFirstName());
}

func (m *myStruct) printFirstName() string{
	return m.FirstName;
}

