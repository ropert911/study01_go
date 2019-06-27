package main

import "fmt"

var (
	FirstName, SecondNames, ThirdNames string
	i                                  int
	f                                  float32
	Input                              = "5.2 / 100 / Golang" //用户自定义变量，便于之后对这个字符串的处理。
	format                             = "%f / %d / %s"
)

func main() {
	fmt.Printf("Please enter your full name: ")
	//Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
	fmt.Scanln(&FirstName, &SecondNames)
	//Scanf与其类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取
	// fmt.Scanf("%s %s", &firstName, &lastName)

	fmt.Printf("Hi %s %s!\n", FirstName, SecondNames)

	//从 Input读，使用format格式
	fmt.Sscanf(Input, format, &f, &i, &ThirdNames)
	fmt.Println("From the Input we read: ", f, i, ThirdNames)
}
