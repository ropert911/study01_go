package main

import "fmt"
import "math"

func foo(a int) {
	defer fmt.Println("foo退出来了")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()

	if a < 0 {
		panic("必须输入大于0的数")
	}

	fmt.Println(a, " 的方根为：", math.Sqrt(float64(a)))

}

func main() {
	var a int
	a = 10
	foo(a)

	fmt.Println("============")
	var b int
	b = -10
	foo(b)

	fmt.Println("该goroutine还可以执行")
}
