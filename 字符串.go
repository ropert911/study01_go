package main

import (
	"fmt"
	"strconv"
)

func strTest1() {
	// 将字符串解析为整数，ParseInt 支持正负号，ParseUint 不支持正负号。
	// base 表示进位制（2 到 36），如果 base 为 0，则根据字符串前缀判断，
	// 前缀 0x 表示 16 进制，前缀 0 表示 8 进制，否则是 10 进制。
	// bitSize 表示结果的位宽（包括符号位），0 表示最大位宽。
	value, err := strconv.ParseInt("4F", 16, 8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	value, err = strconv.ParseInt("FF", 16, 16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	//base:0根据字符串判断
	value, err = strconv.ParseInt("0xFF", 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	fmt.Println(strconv.ParseInt("9", 10, 0))
}

func strTest2() {
	s := "0.12345678901234567890"

	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)          // 0.12345679104328156
	fmt.Println(float32(f), err) // 0.12345679

	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err) // 0.12345678901234568
}

func strTest3() {
	tip1 := "genji is a ninja"
	fmt.Println(len(tip1))

	//string==>[]byte
	str2 := "hello"
	data2 := []byte(str2)
	fmt.Println(data2)

	//[]byte=>string
	str2 = string(data2[:])
	fmt.Println(str2)
}

func strTest4() {
	str := "123456"
	fmt.Println(str[:2])
	fmt.Println(str[2:])
	fmt.Println(str[2:4])
	data2 := []byte(str)
	fmt.Println(data2[:2])
	fmt.Println(data2[2:])
	fmt.Println(data2[2:4])
}

// 示例
func main() {
	//根据字符串解析出整数
	strTest1()

	//根据字符串解析出浮点数
	fmt.Println("=================根据字符串解析出浮点数")
	strTest2()

	//string bytes相互转换
	fmt.Println("=================string bytes相互转换")
	strTest3()

	//string与bytes的切片
	fmt.Println("=================string与bytes的切片")
	strTest4()
}
