package service

import (
	"PWS/util"
	"errors"
	"log"
)

// PWS009 砍价 我的信息
func PWS009(openid, topenid string) (string, error) {

	if openid == topenid {
		return "不能为自己砍价！", errors.New("不能为自己砍价")
	}

	GetConn()
	// 检查是否为该用户砍价过
	var count int
	checkSQL := `select count(*) from bargin_record where openid=? and topenid=? and delflg=0`

	err := db.Get(&count, checkSQL, openid, topenid)
	if err != nil {
		return "砍价失败！", errors.New("数据库查询错误")
	}

	if count != 0 {
		log.Printf("openid: %s 已为openid: %s 砍价过了！", openid, topenid)
		return "您已为该用户砍价过了！", errors.New("您已为该用户砍价过了！")
	}
	// 添加一条砍价记录
	insertSQL := `insert into bargin_record(openid, topenid, createtime) values(?,?,?)`

	sqlReturn, _ := db.Prepare(insertSQL)
	defer sqlReturn.Close()

	_, err2 := sqlReturn.Exec(openid, topenid, util.GenerateTime())
	if err2 != nil {
		log.Printf("数据库插入失败,err: %s", err2)
		return "砍价失败！", errors.New("插入数据库失败")
	}

	// 当砍价达到一定次数，产生金额抵扣券
	checkBarginTimes(topenid)

	return "砍价成功！", nil
}

// 砍价次数实体
type barginCount struct {
	BarginTimes int // 砍价次数
	AutoCards   int // 砍价达到次数赠送抵扣券
}

// 当砍价达到一定次数，产生金额抵扣券
func checkBarginTimes(topenid string) {
	var barginTimes, autoCards, bciid int
	barginCount := barginCount{}
	selectSQL := `` +
		`SELECT count(*) bargintimes, (select count(*) from bargin_record where topenid=? and delflg=0 and recordtype=1) autocards ` +
		`FROM bargin_record ` +
		`WHERE topenid=? AND delflg=0 AND recordtype=0`
		/*
			select GROUP_CONCAT(distinct recordtype ) as  recordtypeid , count(*) as count
			from bargin_record
			where topenid = "oGwQ3uJAi6CwRabkUL8rNfHQrxDQ"  and recordtype in (0,1)  and delflg = 0 GROUP BY recordtype
		*/

		/*
			select count(1) FROM bargin_record where topenid = "oGwQ3uJAi6CwRabkUL8rNfHQrxDQ" AND delflg = 0 group by  recordtype HAVING  recordtype = 0 or recordtype = 1
		*/

	db.Get(&barginCount, selectSQL, topenid)

	barginTimes = barginCount.BarginTimes
	autoCards = barginCount.AutoCards

	// 根据次数产生抵扣券id
	if barginTimes >= 61 && autoCards < 5 {
		bciid = 5
	} else if barginTimes >= 53 && autoCards < 4 {
		bciid = 4
	} else if barginTimes >= 45 && autoCards < 3 {
		bciid = 3
	} else if barginTimes >= 30 && autoCards < 2 {
		bciid = 2
	} else if barginTimes >= 20 && autoCards < 1 {
		bciid = 1
	} else {
		return
	}
	// 产生CDKey
	cdkey := util.RandStr(10)

	insertSQL2 := `` +
		`INSERT INTO bargin_record(openid, topenid, cdkey, bci_id, recordtype, createtime) VALUES(?,?,?)`

	sqlReturn2, _ := db.Prepare(insertSQL2)
	defer sqlReturn2.Close()

	sqlReturn2.Exec(topenid, topenid, cdkey, bciid, 1, util.GenerateTime())
}
