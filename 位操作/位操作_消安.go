package 位操作

import "fmt"

//把浮点数转为 1字节整数，1字节小数
func Translate(value float32) string {
	var flag = 0
	if value < 0 {
		flag = 1
		value = 0 - value
	}

	var temp = uint16(value * 100)
	var valueH = uint8(temp / 100)
	var valueL = uint8(temp % 100)

	//处理整数部分
	var vH = uint8(valueH)
	if flag > 0 {
		vH = vH | 0x80
	}

	var vL = valueL

	return fmt.Sprintf("%02x%02x", vL, vH)
}

func main() {
	var temp float32 = 28.630
	fmt.Println(Translate(temp))
	temp = -28.630
	fmt.Println(Translate(temp))
}
