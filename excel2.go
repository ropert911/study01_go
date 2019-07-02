package main

import (
	"github.com/tealeg/xlsx"
)

func main() {
	file, err := xlsx.OpenFile("format.xlsx")
	if err != nil {
		panic(err)
	}
	first := file.Sheets[0]
	row := first.AddRow()
	row.SetHeightCM(1)
	cell := row.AddCell()
	cell.Value = "1"
	cell = row.AddCell()
	cell.Value = "张三"
	cell = row.AddCell()
	cell.Value = "男"
	cell = row.AddCell()
	cell.Value = "18"

	err = file.Save("file.xlsx")
	if err != nil {
		panic(err)
	}

}
