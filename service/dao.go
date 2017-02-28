package service

import (
	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	mysqlUser    = "linjie"
	mysqlPass    = "linjie"
	mysqlHost    = "localhost"
	mysqlPort    = "3306"
	databaseName = "pwsDB"
)

var (
	db  *sqlx.DB
	err error
)

func reconnect() {
	if db != nil && db.Ping() == nil {
		return
	}
	db, err = sqlx.Open("mysql", mysqlUser+":"+mysqlPass+"@tcp("+mysqlHost+":"+mysqlPort+")/"+databaseName)
	// db, _ := sqlx.Open("mysql", "report:report_1028@tcp(120.26.124.79:3361)/hutong_20161010?charset=utf8")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
}

func init() {
	reconnect()
}

/*
GetConn ...
*/
func GetConn() *sqlx.DB {
	reconnect()
	return db
}
