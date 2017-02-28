package service

import (
	"PWS/util"
	"log"
)

// PWS011 抽奖
func PWS011(openid, topenid string) (int, error) {

	GetConn()
	updateSQL := `` +
		`UPDATE bargin_record SET drawtime=?, cdkey=?, bci_id=? ` +
		`WHERE openid=? and topenid=? AND delflg=0`

	// 产生随机数
	var randomNum int
	switch util.RandNum(5) {
	case 0:
		randomNum = 6
	case 1:
		randomNum = 7
	case 2:
		randomNum = 8
	case 3:
		randomNum = 9
	case 4:
		randomNum = 10
	}
	// 产生CDKey
	cdkey := util.RandStr(10)
	_, err := db.Exec(updateSQL, util.GenerateTime(), cdkey, randomNum, openid, topenid)

	if err != nil {
		log.Printf("数据库更新错误: %s", err)
		return 0, err
	}

	return randomNum, nil
}
