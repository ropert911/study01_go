package main

import (
	"fmt"
	"strconv"
)

/*
	1: Alert Data
			0: Force  0  action force, 1 reaction force
			1: Alert  0 心跳, 1 开盖
			2: 校准   0 初始化成功, 1 初始化失败
			3: 电量   0 电池电量正常, 1 电池电量不足
			4~7:未定义
	2: 电池电压: 将16进制无符号数转为10进制数，然后除以10。
	3. 温度:从 -127 到 +127 将16进制有符号数转为10进制数
	4-5:X 	从 -1023 到+1023  将16进制有符号数转为10进制数
	6-7:Y
	8-9:Z
*/

type SewerCover struct {
	b1Force       uint8 // 0x01
	b1Alert       uint8 // 0x02
	b1Calibration uint8 // 0x04
	b1Voltage     uint8 // 0x08
	b2voltage     uint8 //
	b3temperature int8
	b45SensorX    int16
	b67SensorY    int16
	b89SensorZ    int16
}

func sewerCoverTest() {
	var scData SewerCover
	scData.b1Force = 0
	scData.b1Alert = 0
	scData.b1Calibration = 0
	scData.b1Voltage = 0
	scData.b2voltage = 33
	scData.b3temperature = 27
	scData.b45SensorX = -27
	scData.b67SensorY = -35
	scData.b89SensorZ = 954

	var dataType uint8 = 0
	if scData.b1Force > 0 {
		dataType |= 0x01
	}
	if scData.b1Alert > 0 {
		dataType |= 0x02
	}
	if scData.b1Calibration > 0 {
		dataType |= 0x04
	}
	if scData.b1Voltage > 0 {
		dataType |= 0x08
	}

	fmt.Println(scData)
	data := fmt.Sprintf("%02x%02x%02x%04x%04x%04x", dataType, scData.b2voltage, uint8(scData.b3temperature),
		uint16(scData.b45SensorX), uint16(scData.b67SensorY), uint16(scData.b89SensorZ))
	fmt.Println(data)

	var scData2 SewerCover
	//data="031e19f931fea80408"
	v1, _ := strconv.ParseUint(data[0:2], 16, 16)
	scData2.b1Force = uint8(v1) & 0x01
	scData2.b1Alert = (uint8(v1) & 0x02) >> 1
	scData2.b1Calibration = (uint8(v1) & 0x04) >> 2
	scData2.b1Voltage = (uint8(v1) & 0x08) >> 3
	v1, _ = strconv.ParseUint(data[2:4], 16, 16)
	scData2.b2voltage = uint8(v1)
	v2, _ := strconv.ParseInt(data[4:6], 16, 16)
	scData2.b3temperature = int8(v2)
	v2, _ = strconv.ParseInt(data[6:10], 16, 32)
	scData2.b45SensorX = int16(v2)
	v2, _ = strconv.ParseInt(data[10:14], 16, 32)
	scData2.b67SensorY = int16(v2)
	v2, _ = strconv.ParseInt(data[14:18], 16, 32)
	scData2.b89SensorZ = int16(v2)

	fmt.Println(scData2)
}

func main() {
	sewerCoverTest()
}
