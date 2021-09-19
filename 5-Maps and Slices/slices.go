package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string
	var mySliceInt []int

	mySlice = append(mySlice, "Turkey")
	mySlice = append(mySlice, "China")
	mySlice = append(mySlice, "Germany")
	log.Println(mySlice)

	mySliceInt = append(mySliceInt, 2)
	mySliceInt = append(mySliceInt, 1)
	mySliceInt = append(mySliceInt, 3)

	sort.Ints(mySliceInt)
	log.Println(mySliceInt)

	//shorthand

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	//arrange
	log.Println(numbers[0:2]) // 2.index is not including

	log.Println(numbers[6:9]) //return 6,7,8 value of numbers

	//slice of string
	countries := []string{"Turkey", "China", "Germany"}
	log.Println(countries)
}
