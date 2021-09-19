/*
-Go routine de calısan method return donuyorsa orda hata alırız.Çünkü main go routine calısırken return donen degerin gelip gelmediğini bilemiyoruz.
Bu yuzden main goroutine devam edemez ve programın ilerleyişini return durumuna bagımlı sekilde bırakamayız.
Programın ilerleyişini asla bir return degerine bağlayamayız.(Syntax hatası alırız.)

--Channels : Go routinelerin birbirleriyle iletişime gecmelerini saglar.Aynı zamanda bir Go routine tarafından gonderilen degerin diğer
go routinde main goroutinde kullanılmadan once de gelmesini garanti eder.
*/

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type circle struct {
// 	r float64
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.r * c.r
// }

// func (c circle) display() {
// 	fmt.Println("A Circle")
// }
// func main() {
// 	c1 := circle{5}
// 	go c1.display()
// 	area1 := c1.area()
// 	fmt.Printf("%.2f\n", area1)
// }

//2.Ex:
// package main
// import (
// 	"fmt"
// )

// func merhaba(chan1 chan string) {
// 	// return yerine :
// 	chan1 <- "Merhaba" // Merhabayı chan1 e gonderiyor.Atama değil ama!.
// }
// func main() {

// 	myChannel := make(chan string)
// 	go merhaba(myChannel)

// 	fmt.Println(<-myChannel);

// }

package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")
}
func area(c circle, myChannel chan float64) {
	result := math.Pi * c.r * c.r
	myChannel <- result
}

func main() {
	c1 := circle{5}
	chan1 := make(chan float64)

	go area(c1, chan1)
	fmt.Printf("%2.f\n", <-chan1)
	c1.display();
}
