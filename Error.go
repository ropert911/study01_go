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
	err = fmt.Errorf("未知错误 %s", "123")
	fmt.Println(err)
}
