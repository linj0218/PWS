package service

import (
	"PWS/entity"
	"log"
)

// PWS010 卡券列表
func PWS010(openid string) ([]entity.BarginCard, error) {
	cards := []entity.BarginCard{}

	GetConn()
	selectSQL := "" +
		"SELECT bci.type, bci.`name`, br.cdkey, br.isused, br.usedtime " +
		"FROM bargin_record br LEFT JOIN bargin_card_info bci ON br.bci_id=bci.bci_id " +
		"WHERE br.topenid=? AND br.delflg=0 AND bci.delflg=0"

	db.Select(&cards, selectSQL, openid)
	if err != nil {
		log.Printf("数据库查询错误: %s", err)
		return cards, err
	}

	return cards, nil
}
