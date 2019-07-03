package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func CreateFirstAreaName(topIndex int) string {
	return fmt.Sprintf("PerTest%d", topIndex)
}
func CreateSecAreaName(topIndex int, secIndex int) string {
	return fmt.Sprintf("PerTest%d", topIndex)
}
func CreateSpotAreaName(topIndex int, secIndex int, spotIndex int) string {
	return fmt.Sprintf("PerTest%d", topIndex)
}

type ITFInfo struct {
	top           string //一级区域
	sec           string //二级区域
	spot          string //三级区域
	detailAddress string //详细地址
	itfName       string //ITF名称
	itfGWIp       string //ITF网关IP
	itfLng        string //经度
	itfLat        string //纬度
	comment       string //备注
}

func AddITFRow(sheet *xlsx.Sheet, itfInfo ITFInfo) {
	row := sheet.AddRow()
	//row.SetHeightCM(1) //设置每行的高度
	row.AddCell().SetValue(itfInfo.top)
	row.AddCell().SetValue(itfInfo.sec)
	row.AddCell().SetValue(itfInfo.spot)
	row.AddCell().SetValue(itfInfo.detailAddress)
	row.AddCell().SetValue(itfInfo.itfName)
	row.AddCell().SetValue(itfInfo.itfGWIp)
	row.AddCell().SetValue(itfInfo.itfLng)
	row.AddCell().SetValue(itfInfo.itfLat)
	row.AddCell().SetValue(itfInfo.comment)
}

type DevInfo struct {
	top               string      //一级区域
	sec               string      //二级区域
	spot              string      //三级区域
	detailAddress     string      //详细地址
	devID             string      //设备ID
	devName           string      //设备名称
	devIdentification string      //设备标识
	devManufacturer   string      //设备厂家
	devType           string      //设备类型
	devVersion        string      //设备版本
	devModel          string      //设备型号
	gwUpProtocol      string      //网关上行协议
	devLng            interface{} //经度
	devLat            interface{} //纬度
	heartTime         interface{} //心跳周期
	carNum            string      //车牌号
	comment           string      //备注
}

func AddSensorRow(sheet *xlsx.Sheet, info DevInfo) {
	row := sheet.AddRow()
	row.AddCell().SetValue(info.top)
	row.AddCell().SetValue(info.sec)
	row.AddCell().SetValue(info.spot)
	row.AddCell().SetValue(info.detailAddress)
	row.AddCell().SetValue(info.devID)
	row.AddCell().SetValue(info.devName)
	row.AddCell().SetValue(info.devIdentification)
	row.AddCell().SetValue(info.devManufacturer)
	row.AddCell().SetValue(info.devType)
	row.AddCell().SetValue(info.devVersion)
	row.AddCell().SetValue(info.devModel)
	row.AddCell().SetValue(info.gwUpProtocol)
	row.AddCell().SetValue(info.devLng)
	row.AddCell().SetValue(info.devLat)
	row.AddCell().SetValue(info.heartTime)
	row.AddCell().SetValue(info.carNum)
	row.AddCell().SetValue(info.comment)
}

func CreateGwInfo(topIndex int, secIndex int, spotIndex int) DevInfo {
	sensor := DevInfo{detailAddress: "软件园A1", devManufacturer: "博频", devType: "设备网关", devVersion: "v1.1", gwUpProtocol: "MQTT", devModel: "WLRGFM-100"}
	sensor.top = CreateFirstAreaName(topIndex)
	sensor.sec = CreateSecAreaName(topIndex, secIndex)
	sensor.spot = CreateSpotAreaName(topIndex, secIndex, spotIndex)
	sensor.devID = fmt.Sprintf("01%04x%04x%04x00\n", topIndex, secIndex, spotIndex)
	sensor.devName = fmt.Sprintf("LoRa WAN 网关-TEST%d_%d_%d", topIndex, secIndex, spotIndex)
	sensor.devIdentification = fmt.Sprintf("01%04x%04x%04x00\n", topIndex, secIndex, spotIndex)
	sensor.devLng = 104.070313
	sensor.devLat = 30.550611
	sensor.heartTime = 86500
	sensor.comment = fmt.Sprintf("GW Comment%d_%d_%d", topIndex, secIndex, spotIndex)

	return sensor
}

func CreateCommonSaferInfo(topIndex int, secIndex int, spotIndex int) DevInfo {
	sensor := DevInfo{detailAddress: "软件园A1", devManufacturer: "消安", devVersion: "v1.1", gwUpProtocol: "MQTT"}
	sensor.top = CreateFirstAreaName(topIndex)
	sensor.sec = CreateSecAreaName(topIndex, secIndex)
	sensor.spot = CreateSpotAreaName(topIndex, secIndex, spotIndex)
	return sensor
}
func CreateSaferOtherInfo(sensor DevInfo, topIndex int, secIndex int, spotIndex int, devIndex int, devType string, devModel string, devLng float64, devLat float64, heartTime int) DevInfo {
	sensor.devID = fmt.Sprintf("02%04x%04x%04x%02x\n", topIndex, secIndex, spotIndex, devIndex)
	sensor.devName = fmt.Sprintf("%s%04x%04x%04x", devModel, topIndex, secIndex, spotIndex)
	sensor.devIdentification = fmt.Sprintf("GLN%04x%04x%04x%02x", topIndex, secIndex, spotIndex, devIndex)
	sensor.devType = devType
	sensor.devModel = devModel
	sensor.devLng = devLng
	sensor.devLat = devLat
	sensor.comment = fmt.Sprintf("%s Commont %d_%d_%d", devModel, topIndex, secIndex, spotIndex)
	sensor.heartTime = heartTime
	return sensor
}
func CreateSchoolBusInfo(topIndex int, secIndex int, spotIndex int, devIndex int) DevInfo {
	sensor := DevInfo{devManufacturer: "微思格", devType: "校车", devVersion: "v1.1", devModel: "xxx", gwUpProtocol: "MQTT", devLng: 104.070363, devLat: 30.547711, heartTime: 3600}
	sensor.top = CreateFirstAreaName(topIndex)
	sensor.sec = CreateSecAreaName(topIndex, secIndex)
	sensor.spot = CreateSpotAreaName(topIndex, secIndex, spotIndex)
	sensor.detailAddress = "软件园A1"
	sensor.devID = fmt.Sprintf("03%04x%04x%04x%02x\n", topIndex, secIndex, spotIndex, devIndex)
	sensor.devName = "校车" + sensor.devID
	sensor.carNum = fmt.Sprintf("川%d%d%d%d", topIndex, secIndex, spotIndex, devIndex)
	sensor.devIdentification = sensor.carNum
	sensor.comment = fmt.Sprintf("校车%d_%d_%d %d", topIndex, secIndex, spotIndex, devIndex)
	return sensor
}

func CreateSweatCoverInfo(topIndex int, secIndex int, spotIndex int, devIndex int) DevInfo {
	sensor := DevInfo{devIdentification: "WSMS-125", devManufacturer: "博频", devType: "智能井盖", devVersion: "v1.1", devModel: "WSMS-125", gwUpProtocol: "MQTT", devLng: 104.069347, devLat: 30.546741, heartTime: 3600}
	sensor.top = CreateFirstAreaName(topIndex)
	sensor.sec = CreateSecAreaName(topIndex, secIndex)
	sensor.spot = CreateSpotAreaName(topIndex, secIndex, spotIndex)
	sensor.detailAddress = "软件园A1"
	sensor.devID = fmt.Sprintf("04%04x%04x%04x%02x\n", topIndex, secIndex, spotIndex, devIndex)
	sensor.devName = fmt.Sprintf("智能井盖%x%x%x%d", topIndex, secIndex, spotIndex, devIndex)
	sensor.comment = fmt.Sprintf("井盖%d_%d_%d %d", topIndex, secIndex, spotIndex, devIndex)

	return sensor
}
func CreateSmokeInfo(topIndex int, secIndex int, spotIndex int, devIndex int) DevInfo {
	sensor := DevInfo{devIdentification: "GS530L", devManufacturer: "烟感", devType: "烟感报警器", devVersion: "v1.1", devModel: "GS530L", gwUpProtocol: "MQTT", devLng: 104.071348, devLat: 30.548731, heartTime: 3600}
	sensor.top = CreateFirstAreaName(topIndex)
	sensor.sec = CreateSecAreaName(topIndex, secIndex)
	sensor.spot = CreateSpotAreaName(topIndex, secIndex, spotIndex)
	sensor.detailAddress = "软件园A1"
	sensor.devID = fmt.Sprintf("04%04x%04x%04x%02x\n", topIndex, secIndex, spotIndex, devIndex)
	sensor.devName = fmt.Sprintf("烟感报警器%x%x%x%d", topIndex, secIndex, spotIndex, devIndex)
	sensor.comment = fmt.Sprintf("烟感%d_%d_%d %d", topIndex, secIndex, spotIndex, devIndex)
	return sensor
}
func AddThirdAreaSensors(sheet *xlsx.Sheet, topIndex int, secIndex int, spotIndex int) {
	//添加GateWay
	AddSensorRow(sheet, CreateGwInfo(topIndex, secIndex, spotIndex))

	//消安传感器
	sensor := CreateCommonSaferInfo(topIndex, secIndex, spotIndex)
	AddSensorRow(sheet, CreateSaferOtherInfo(sensor, topIndex, secIndex, spotIndex, 1, "水位传感器-WDS", "WDS", 104.070163, 30.550011, 3600))
	AddSensorRow(sheet, CreateSaferOtherInfo(sensor, topIndex, secIndex, spotIndex, 2, "水压传感器-WPS", "WPS", 104.070213, 30.550201, 3600))
	AddSensorRow(sheet, CreateSaferOtherInfo(sensor, topIndex, secIndex, spotIndex, 3, "温湿度传感器-THS", "THS", 104.071323, 30.551721, 3600))
	AddSensorRow(sheet, CreateSaferOtherInfo(sensor, topIndex, secIndex, spotIndex, 4, "开关量传感器-SNO", "SNO", 104.069023, 30.549921, 3600))
	AddSensorRow(sheet, CreateSaferOtherInfo(sensor, topIndex, secIndex, spotIndex, 5, "三相交流电压传感器-PVM", "PVM", 104.069133, 30.549831, 3600))
	AddSensorRow(sheet, CreateSaferOtherInfo(sensor, topIndex, secIndex, spotIndex, 6, "三相交流电流传感器-PCM-L", "PCM-L", 104.070643, 30.549741, 3600))
	AddSensorRow(sheet, CreateSaferOtherInfo(sensor, topIndex, secIndex, spotIndex, 7, "三相交流电流传感器-PCM-H", "PCM-H", 104.070745, 30.549043, 3600))
	AddSensorRow(sheet, CreateSaferOtherInfo(sensor, topIndex, secIndex, spotIndex, 8, "剩余电流传感器-PRC", "PRC", 104.070845, 30.550843, 3600))
	AddSensorRow(sheet, CreateSaferOtherInfo(sensor, topIndex, secIndex, spotIndex, 9, "三相交流电源开关状态传感器-PSD", "PSD", 104.070946, 30.550946, 3600))

	AddSensorRow(sheet, CreateSchoolBusInfo(topIndex, secIndex, spotIndex, 10))
	//井盖
	AddSensorRow(sheet, CreateSweatCoverInfo(topIndex, secIndex, spotIndex, 11))
	AddSensorRow(sheet, CreateSmokeInfo(topIndex, secIndex, spotIndex, 12))

}
func main() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("ITF")
	if err != nil {
		fmt.Println(err)
		return
	}

	AddITFRow(sheet, ITFInfo{top: "一级区域", sec: "二级区域", spot: "三级区域", detailAddress: "详细地址", itfName: "ITF名称", itfGWIp: "ITF网关IP", itfLng: "经度", itfLat: "纬度", comment: "备注"})
	AddITFRow(sheet, ITFInfo{top: "PerTest0", sec: "PerTest0_0", spot: "PerTest0_0_0", detailAddress: "软件园A1", itfName: "20.47", itfGWIp: "192.168.20.47", itfLng: "104.070163", itfLat: "30.550011"})

	sheet, err = file.AddSheet("GW-SENSOR")
	if err != nil {
		fmt.Println(err)
		return
	}

	AddSensorRow(sheet, DevInfo{top: "一级区域", sec: "二级区域", spot: "三级区域", detailAddress: "详细地址", devID: "设备ID", devName: "设备名称", devIdentification: "设备标识", devManufacturer: "设备厂家", devType: "设备类型", devVersion: "设备版本", devModel: "设备型号", gwUpProtocol: "网关上行协议", devLng: "经度", devLat: "纬度", heartTime: "心跳周期", carNum: "车牌号", comment: "备注"})
	AddThirdAreaSensors(sheet, 1, 1, 1)
	AddThirdAreaSensors(sheet, 1, 1, 2)

	err = file.Save("C:\\Users\\sk-qianxiao\\Desktop\\区域.xlsx")
	if err != nil {
		panic(err)
	}

}
