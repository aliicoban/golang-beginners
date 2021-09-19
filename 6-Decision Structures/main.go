package main

import "log"

func main() {
	myNum := 100
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	if myNum > 100 && isTrue {
		log.Println("Mynum greater than 100 and isTrue is true")
	} else if myNum < 100 || isTrue {
		log.Println("mynum greater than 100 or isTrue is true")
	} else {
		log.Println("Other cases")
	}

}
