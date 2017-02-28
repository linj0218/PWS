package service

import (
	"PWS/entity"
	"PWS/util"
	"fmt"
)

type info struct {
	A string `json:"a"`
	B string `json:"b"`
}

// PWS002 读取info.xlsx文件并返回数据
func PWS002() entity.ResEntity {
	fileContent := util.ReadExcel("./files/info_v3.xlsx")

	infoList := []info{}

	for _, row := range fileContent[0] {
		temp := info{A: row[0], B: row[1]}
		infoList = append(infoList, temp)
	}
	fmt.Println("infoList: ", infoList)

	resp := entity.ResEntity{}
	resp.Status = 0
	resp.Msg = "成功"
	resp.Data = infoList

	return resp
}
