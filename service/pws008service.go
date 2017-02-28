package service

import (
	"PWS/entity"
	"log"
)

// PWS008 砍价 我的信息
func PWS008(openid string) (entity.BarginUserInfo, error) {

	userInfo := entity.BarginUserInfo{}

	GetConn()
	selectSQL := `` +
		`SELECT ` +
		`wui.nickname, wui.headimgurl, (SELECT count(*) FROM bargin_record br WHERE br.topenid=? AND br.delflg=0) bbargintime ` +
		`FROM wx_user_info wui WHERE wui.openid=? AND wui.delflg=0`

	err := db.Get(&userInfo, selectSQL, openid, openid)

	if err != nil {
		log.Printf("数据查询错误: %s", err)
		return userInfo, err
	}

	userInfo.TargetTime = "2017-02-01 00:00:00"

	return userInfo, nil
}
