/*
-Go routine ler ile birden fazla işin aynı anda yapılmasını sağlayabilirz.(Eşzamanlı)
-Bir programın en hızlı calısma sekli sequence değil , sıralı sekilde calısma en verimli değil.
-Main fonks calısmaya basladıgı anda main go routine olusturulur ve calısmaya baslar. 
-Goroutinelerin birbirleri arasında ne zaman ve nasıl gecis yapacagını ve sürelerini sadece goroutine yazarak bilemiyoruz.
-Tek cekirdekli sistemlerde bir işlemci üzerinde o an üzerinde tek bir gorutine calısır.Ama gecikme olursa diğer goroutineler devreye girebilir.
Go düzenleyici o anda devreye hangisini alacagını en effective sekilde belirtiyor.
*/

package main

import "fmt"

type shape interface {
	area() float64
}

func printArea(shapes ...shape) {
	for _, shape := range shapes {
		fmt.Println("Area is : ", shape.area())
	}
}

type triangle struct {
	a float64
	h float64
}

func (t triangle) area() float64 {
	return (t.a * t.h) / 2
}

type square struct {
	a float64
}

func (s square) area() float64 {
	return (s.a * s.a)
}

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {
	return (r.a * r.b)
}

func main() {

	t := triangle{3, 4}
	s := square{4}
	r := rectangle{4, 5}

	printArea(t, s, r)

}