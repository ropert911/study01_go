package main

//本示例有2个真实线程；现个协程可同时执行

import (
	"fmt"
	"runtime"
)

func loop(done chan bool) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", i)
		//runtime.Gosched()  	// 显式地让出CPU时间给其他goroutine，这种一般不用
	}
	done <- true
}

func main() {
	runtime.GOMAXPROCS(2) //设置真实线程数
	done := make(chan bool)
	go loop(done)
	go loop(done)

	<-done
	<-done

}
