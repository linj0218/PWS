package util

import (
	"fmt"

	"strconv"

	"strings"

	"github.com/tealeg/xlsx"
)

/*
ReadExcel 读excel文件
*/
func ReadExcel(filepath string) [][][]string {
	var dataList [][][]string
	excelFileName := filepath
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("err: ", err)
		fmt.Println("create a new one")
		file := xlsx.NewFile()
		_, err := file.AddSheet("Sheet1")
		if err != nil {
			fmt.Printf(err.Error())
		}
		err = file.Save(filepath)
		if err != nil {
			fmt.Printf(err.Error())
		}
		return nil
	}
	// sheets
	var sheetList [][][]string
	for _, sheet := range xlFile.Sheets {
		// rows
		var rowList [][]string
		for _, row := range sheet.Rows {
			// cells
			var cellList []string
			for _, cell := range row.Cells {
				t, _ := cell.String()
				if t == "" {
					continue
				}
				cellList = append(cellList, t)
				// fmt.Printf("%s\n", t)
			}
			rowList = append(rowList, cellList)
		}
		sheetList = append(sheetList, rowList)
	}
	dataList = sheetList
	fmt.Println("读取结果: ", dataList)
	return dataList
}

/*
WriteExcel 写excel文件
*/
func WriteExcel(filepath string, dataList [][][]string) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	for n, x := range dataList {
		sheetName := []string{"Sheet", strconv.Itoa(n)}
		sheet, err = file.AddSheet(strings.Join(sheetName, ""))
		if err != nil {
			fmt.Printf(err.Error())
		}
		for _, y := range x {
			row = sheet.AddRow()
			for _, z := range y {
				cell = row.AddCell()
				cell.Value = z
			}
		}
	}
	err = file.Save(filepath)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
