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

// 示例
func main() {
	//根据字符串解析出整数
	//strTest1()
	//浮点数解析
	strTest2()
}
