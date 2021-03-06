package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // waiting operations

func main() {

	wg.Add(1)

	go printX()
	wg.Wait()
	go printY()

}

func printX() {
	for i := 0; i < 1000; i++ {
		fmt.Print("X")
	}
	wg.Done()
}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
}