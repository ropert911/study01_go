package main

import "fmt"

func translate(value float32) string {
	var flag = 0
	if value < 0 {
		flag = 1
		value = 0 - value
	}

	var valueL = value - float32(int(value))
	var valueH = value - valueL

	//处理整数部分
	var vH = uint8(valueH)
	if flag > 0 {
		vH = vH | 0xf0
	}

	//处理小数部分
	var vL uint8 = 0
	for i := 0; i < 8; i++ {
		vL = vL << 1
		valueL = valueL * 2
		if valueL >= 1 {
			valueL = valueL - 1
			vL = vL | 0x1
		}
	}
	return fmt.Sprintf("%02x%02x", vH, vL)
}
func main() {
	//fmt.Println(fmt.Sprintf("%s://%s:%v", "tcp", "1.2.2.3", ""))
	//fmt.Println("get address:", "aaaaaaaaa")
	var value float32 = 13.125
	fmt.Println(translate(value))
	value = -13.125
	fmt.Println(translate(value))
}
