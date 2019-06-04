package main

import (
	"fmt"
)

const (
	version     string = "1.2.3"
	serviceName string = "device-random"
)

func main() {

	fmt.Println(serviceName, version)
}
