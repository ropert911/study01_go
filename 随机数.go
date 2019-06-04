package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(10)
	fmt.Println(a)
	a = rand.Intn(10)
	fmt.Println(a)
	a = rand.Intn(10)
	fmt.Println(a)
}
