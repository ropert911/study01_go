package main

import "fmt"

func test2() {
	a := "xyz"
	b := "opq"
	pa := &a  //pa为指向a的指针
	pp := &pa //pp为指向pa的指针
	fmt.Println(a, b, *pa, **pp)
	a += "zz" //a追加“zz”
	fmt.Println(a, b, *pa, **pp)
	*pa += "bb" //pp指向的值，追加"bb"
	fmt.Println(a, b, *pa, **pp)
	fmt.Println("打印a各种情况：", &a, a)
	fmt.Println("打印pa各种情况：", &pa, pa, *pa)
	fmt.Println("打印pp各种情况：", &pp, pp, *pp, **pp)
}

func main() {
	test2()
}
