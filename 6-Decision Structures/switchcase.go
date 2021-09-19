package main

import "log"

func main() {
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("Cat is set to cat")
	case "dog":
		log.Println("Cat is set to dog")
	case "fish":
		log.Println("Cat is set to fish")
	default:
		log.Println("Cat is something else")
	}
}
