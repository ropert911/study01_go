package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func call() {
	fmt.Println("aaaaaaaaaaaa")
}

func main() {
	once.Do(call)
	once.Do(call)
	fmt.Println("exit")
}
