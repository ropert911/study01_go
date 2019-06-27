package main

import (
	"fmt"
	"time"
)

func doneSelect(ch chan int) {
	for {
		//fmt.Println("for循环-输出-start")
		select {
		case data := <-ch:
			fmt.Println(data)
			break // 使用break只会跳出select
		default:
			//fmt.Println("select default testing")
			break
		}
	}
	fmt.Println("for循环-跳出-end **********")
}
func do() {
	ch := make(chan int)
	go doneSelect(ch)
	ch <- 666
}
func main() {
	do()
	time.Sleep(3 * time.Second)
}
