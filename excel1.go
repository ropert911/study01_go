package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func AddITFHeader(sheet *xlsx.Sheet, top string, sec string, sport string, address string, itfName string, itfGWip string, lng string, lat string, comment string) {
	row := sheet.AddRow()
	//row.SetHeightCM(1) //设置每行的高度
	cellFirst := row.AddCell()
	cellSec := row.AddCell()
	cellSpot := row.AddCell()
	cellDetailedAddress := row.AddCell()
	cellITFName := row.AddCell()
	cellITFGWIP := row.AddCell()
	cellLng := row.AddCell()
	cellLat := row.AddCell()
	cellComment := row.AddCell()

	cellFirst.Value = top
	cellSec.Value = sec
	cellSpot.Value = sport
	cellDetailedAddress.Value = address
	cellITFName.Value = itfName
	cellITFGWIP.Value = itfGWip
	cellLng.Value = lng
	cellLat.Value = lat
	cellComment.Value = comment
}
func main() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("ITF")
	if err != nil {
		fmt.Println(err)
		return
	}

	AddITFHeader(sheet, "一级区域", "二级区域", "三级区域", "详细地址", "ITF名称", "ITF网关IP", "经度", "纬度", "备注")
	AddITFHeader(sheet, "PerTest0", "PerTest0_0", "PerTest0_0_0", "软件园A1", "20.47", "192.168.20.47", "104.070163", "30.550011", "")

	err = file.Save("C:\\Users\\sk-qianxiao\\Desktop\\区域.xlsx")
	if err != nil {
		panic(err)
	}

}
