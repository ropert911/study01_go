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
	b1Force       byte    // 0x80
	b1Alert       byte    // 0x40
	b1Calibration byte    // 0x20
	b1Voltage     byte    // 0x10
	b2voltage     float32 //
	b3temperature byte
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
	scData.b2voltage = 3.3
	scData.b3temperature = 27
	scData.b45SensorX = -27
	scData.b67SensorY = -35
	scData.b89SensorZ = 954

	var dataType byte = 0
	if scData.b1Force > 0 {
		dataType |= 0x80
	}
	if scData.b1Alert > 0 {
		dataType |= 0x40
	}
	if scData.b1Calibration > 0 {
		dataType |= 0x20
	}
	if scData.b1Voltage > 0 {
		dataType |= 0x10
	}

	fmt.Println(scData)
	data := fmt.Sprintf("%02x%02x%02x%04x%04x%04x", dataType, byte(scData.b2voltage*10), scData.b3temperature,
		uint16(scData.b45SensorX), uint16(scData.b67SensorY), uint16(scData.b89SensorZ))
	fmt.Println(data)
	fmt.Println()

	var scData2 SewerCover
	v1, _ := strconv.ParseInt(data[0:2], 16, 8)
	scData2.b1Force = byte(v1) & 0x80
	scData2.b1Alert = byte(v1) & 0x40
	scData2.b1Calibration = byte(v1) & 0x20
	scData2.b1Voltage = byte(v1) & 0x10
	v1, _ = strconv.ParseInt(data[2:4], 16, 16)
	scData2.b2voltage = float32(v1) / 10
	v1, _ = strconv.ParseInt(data[4:6], 16, 16)
	scData2.b3temperature = byte(v1)
	v1, _ = strconv.ParseInt(data[6:10], 16, 32)
	scData2.b45SensorX = int16(v1)
	v1, _ = strconv.ParseInt(data[10:14], 16, 32)
	scData2.b67SensorY = int16(v1)
	v1, _ = strconv.ParseInt(data[14:18], 16, 32)
	scData2.b89SensorZ = int16(v1)

	fmt.Println(scData2)
}

func main() {

}
