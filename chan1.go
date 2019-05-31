package main

import (
	"fmt"
	"time"
)

func Product(ch chan<- int) {
	for i := 0; i < 100; i++ {
		fmt.Println("Product:", i)
		ch <- i
	}
}
func Consumer(ch <-chan int) {
	for i := 0; i < 100; i++ {
		a := <-ch
		fmt.Println("Consmuer:", a)
	}
}
func main() {
	ch := make(chan int, 0) //大小为零，就写了就阻塞；大小为1，可写1个，下次再写就阻塞
	go Product(ch)
	go Consumer(ch)
	time.Sleep(50)
}
