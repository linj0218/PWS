package service

import (
	"fmt"
	"testGo/util"
)

// InsertUser 测试插入数据
func InsertUser() {
	GetConn()
	insertSQL := `insert into userinfo(name,pass,sex,tel,addr,idNum,createTime) values(?,?,?,?,?,?,?)`
	sqlReturn, err := db.Prepare(insertSQL)
	defer sqlReturn.Close()
	_, err2 := sqlReturn.Exec("小六", "xiaoliu", "男", "12345678914", "上海市", "123456789012345674", util.GenerateTime())
	if err != nil || err2 != nil {
		fmt.Println("数据库插入失败")
	}
}

type userInfo struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Tel   string `json:"tel"`
	Addr  string `json:"addr"`
	Idnum string `json:"idNum"`
}

// SelectUserById 按id查询用户
func SelectUserById(id int) []userInfo {
	GetConn()
	userInfo := []userInfo{}
	selectSQL := `select id,name,sex,tel,addr,idnum from userinfo where id=?`
	err := db.Select(&userInfo, selectSQL, id)
	if err != nil {
		fmt.Println("数据库查询错误")
	}
	return userInfo
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(id int, name, pass, sex, tel, addr, idnum string) {
	GetConn()
	updateSQL := `update userinfo set name=?,pass=?,sex=?,tel=?,addr=?,idnum=? where id=?`
	_, err := db.Exec(updateSQL, name, pass, sex, tel, addr, idnum, id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("数据库更新错误")
	}

	// rows1, err1 := result1.RowsAffected()
	// if err1 == nil {
	// 	if rows1 > 0 {

	// 	} else {
	// 		return wxResponse.ConfigError(1, "删除失败")
	// 	}
	// }

}

// DeleteUserById 按id删除用户
func DeleteUserById(id int) {
	deleteSQL := `update userinfo set delflg=1 where id=?`
	_, err := db.Exec(deleteSQL, id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("数据库删除错误")
	}
}
