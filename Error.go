package main

import (
	"errors"
	"fmt"
)

func SomeFunc() error {
	if true {
		return errors.New("遇到了某某错误")
	}
	return nil
}

func main() {
	err := SomeFunc()
	fmt.Println(err)
}
