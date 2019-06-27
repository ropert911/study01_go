package main

import (
	"fmt"
	"os"
)

const Version string = "0.7.0"
const (
	version            = Version
	serviceName string = "device-random"
)

type Service struct {
	name string
}

var (
	svc Service
)

func main() {

	fmt.Println(serviceName, version)

	err := fmt.Errorf("NewService: empty name specified\n")
	fmt.Println("11111111")
	fmt.Fprintf(os.Stderr, "error: %v", err)
	fmt.Println("2222222")

	svc.name = "server name"
	fmt.Println(svc)

	os.Exit(1)
	fmt.Println("The end print")
}
