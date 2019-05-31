package main

//本示例只有一个线程，两个协和；但在执行时，一个协和没有阻塞情况是不会把cpu让给其它程序。
//所以结果是一个跑完了，再跑另一个

import (
	"fmt"
	"os"
	"runtime"
)

func loop(done chan bool) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", i)
	}
	done <- true
}

func main() {
	fmt.Println("当前进程Id:", os.Getpid())
	fmt.Println("cpu数量：", runtime.NumCPU())
	fmt.Println("routin数：", runtime.NumGoroutine())

	done := make(chan bool)
	go loop(done)
	go loop(done)

	<-done
	<-done

}
