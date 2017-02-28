package util

import "time"

/*
GenerateDate generate and returns the current date.
return "2006-01-02"
*/
func GenerateDate() string {
	timestamp := time.Now().Unix()
	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	return string(tm.Format("2006-01-02"))
}

/*
GenerateTime generate and returns the current time.
return "20060102150405"
*/
func GenerateTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return string(tm.Format("20060102150405"))
}

/*
GenerateFormatTime generate and returns the current time.
return "2006-01-02 15:04:05"
*/
func GenerateFormatTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return string(tm.Format("2006-01-02 15:04:05"))
}
