package main

import (
	"log"
)

type Animal interface {
	Says() string
	NumberofLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main(){
	dog := Dog{
		Name:  "Karabas",
		Breed: "Sivas Kangali",
	}
	PrintInfo(dog);

	gorilla := Gorilla {
		Name :"King Kong",
		Color : "Black",
		NumberOfTeeth: 32,
	}
	PrintInfo(gorilla);
}

func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberofLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "grunt"
}

func (g Gorilla) NumberofLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberofLegs(),"legs");
}
