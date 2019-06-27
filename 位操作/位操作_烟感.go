package 位操作

import (
	"fmt"
	"strconv"
)

//博频-烟感 版本号:v1.03
type SmokeSensor struct {
	b1devType    uint8 //0xf0	1表示为烟雾报警器
	b1sofVersion uint8 //0x0f   协议版本 当前3
	//设备状态：
	//BIT7:保留，默认为0；
	//BIT6:保留，默认为0；
	//BIT5:等于1表示本地测试报警(当传感器故障时按测试键不置位此标记位)；
	//BIT4:等于1表示报警记忆；
	//BIT3:等于1表示烟雾报警器电池电压过低；
	//BIT2:等于1表示烟雾报警器处于传感器故障；
	//BIT1:等于1表示烟雾报警器处于报警静音；
	//BIT0:等于1表示烟雾报警器侦测到烟雾进入报警；
	b2devStatus uint8

	//当前报警器电池电压百分比
	b3volPercentage uint8
	//当前报警器烟雾浓度百分比
	b4smokePercentage uint8
	//报警器当前蜂鸣器的状态  0x01:蜂鸣器处于静音   0x00:蜂鸣器处于正常(报警时能发出报警声音)
	b5Status uint8
	//和校验（BIT7-BIT0）=以上字节数据相加结果为0x00
	b6checkSum uint8
}

func smokeTest() {
	var scData SmokeSensor
	scData.b1devType = 1
	scData.b1sofVersion = 3
	scData.b2devStatus = 0x20 //测试消息
	scData.b3volPercentage = 90
	scData.b4smokePercentage = 2
	scData.b5Status = 0
	var b1 uint8
	b1 = scData.b1sofVersion | (scData.b1devType << 4)
	var b2 uint8
	b2 = scData.b2devStatus
	var sum uint8
	sum = b1 + b2 + scData.b3volPercentage + scData.b4smokePercentage + scData.b5Status
	scData.b6checkSum = ^sum + 1

	fmt.Println(scData)

	var data = fmt.Sprintf("%02x%02x%02x%02x%02x%02x", b1, b2, scData.b3volPercentage, scData.b4smokePercentage, scData.b5Status, scData.b6checkSum)
	fmt.Println(data)

	{
		var smokeSensor SmokeSensor
		var b uint8
		//data = "132064000069"
		v1, _ := strconv.ParseUint(data[0:2], 16, 16)
		b = uint8(v1)
		smokeSensor.b1sofVersion = b & 0x0f
		smokeSensor.b1devType = (b & 0xf0) >> 4
		v1, _ = strconv.ParseUint(data[2:4], 16, 16)
		b = uint8(v1)
		smokeSensor.b2devStatus = b
		v1, _ = strconv.ParseUint(data[4:6], 16, 16)
		smokeSensor.b3volPercentage = uint8(v1)
		v1, _ = strconv.ParseUint(data[6:8], 16, 16)
		smokeSensor.b4smokePercentage = uint8(v1)
		v1, _ = strconv.ParseUint(data[8:10], 16, 16)
		smokeSensor.b5Status = uint8(v1)
		v1, _ = strconv.ParseUint(data[10:12], 16, 16)
		smokeSensor.b6checkSum = uint8(v1)
		fmt.Println(smokeSensor)
	}
}

func main() {
	smokeTest()
}
