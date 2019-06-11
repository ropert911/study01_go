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

type SmokeSensor struct {
	b1aggVersion uint8 //0x0f	协议版本 当前2
	b1sofVersion uint8 //0xf0  软件版本 当前2
	//0x0f  消息类型   1.报警  2.报警静音(静音时间小于100 秒)  3.故障   4.低压  7.正常待机
	// 测试：当松开测试按键时RF 上传报警器状态，例如当前烟感报警器为电池低压时，按测试键将上传0x14； 其他保留  心跳每天一次（心跳上传消息类型为报警器当前的状态）
	b2mesType uint8
	//0xf0  设备类型  1.烟感  2.热感-保留   3.一氧化碳-保留 4.可燃气体 5.温度+烟雾复合型报警器，保留（未使用）
	//6.一氧化碳+烟雾复合型报警器，保留（未使用）
	//7.一氧化碳+烟雾+温度复合型报警器，保留（未使用）；
	//8.一氧化碳+可燃气体复合型报警器，保留（未使用）；
	//9.PIR 探测器，保留（未使用）；
	//A.漏水报警器，保留（未使用）；
	//B.门磁探测器，保留（未使用）；
	//C.温湿探测器，保留（未使用）；
	//其他保留
	b2devType uint8
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
	var smokeSensor SmokeSensor
	smokeSensor.b1aggVersion = 2
	smokeSensor.b1sofVersion = 2
	smokeSensor.b2mesType = 7
	smokeSensor.b2devType = 1
	smokeSensor.b3volPercentage = 90
	smokeSensor.b4smokePercentage = 2
	smokeSensor.b5Status = 0
	var b1 uint8
	b1 = smokeSensor.b1aggVersion | (smokeSensor.b1sofVersion << 4)
	var b2 uint8
	b2 = smokeSensor.b2mesType | (smokeSensor.b2devType << 4)
	var sum uint8
	sum = b1 + b2 + smokeSensor.b3volPercentage + smokeSensor.b4smokePercentage + smokeSensor.b5Status
	smokeSensor.b6checkSum = ^sum + 1

	fmt.Println(smokeSensor)

	data := fmt.Sprintf("%02x%02x%02x%02x%02x%02x", b1, b2, smokeSensor.b3volPercentage, smokeSensor.b4smokePercentage, smokeSensor.b5Status, smokeSensor.b6checkSum)
	fmt.Println(data)

	{
		var smokeSensor SmokeSensor
		var b uint8
		data = "132064000069"
		v1, _ := strconv.ParseUint(data[0:2], 16, 16)
		b = uint8(v1)
		smokeSensor.b1aggVersion = b & 0x0f
		smokeSensor.b1sofVersion = (b & 0xf0) >> 4
		v1, _ = strconv.ParseUint(data[2:4], 16, 16)
		b = uint8(v1)
		smokeSensor.b2mesType = b & 0x0f
		smokeSensor.b2devType = (b & 0xf0) >> 4
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
	//sewerCoverTest()
	smokeTest()
}
