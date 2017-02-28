package service

import (
	"PWS/entity"
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// PWS005 生成jssdk签名
func PWS005(url string) entity.JsapiTicket {
	jsapiticket := GetJsAPITicket()
	noncestr := "Wm3WZYTPz0wzccnW"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	u := url

	// 步骤1. 对所有待签名参数按照字段名的ASCII 码从小到大排序（字典序）后，使用URL键值对的格式（即key1=value1&key2=value2…）拼接成字符串string1
	string1 := strings.Join([]string{"jsapi_ticket=", jsapiticket,
		"&noncestr=", noncestr,
		"&timestamp=", timestamp,
		"&url=", u}, "")

	jsapiTick := entity.JsapiTicket{}

	jsapiTick.AppID = "wx3bd49c0c948957c8"
	jsapiTick.Noncestr = noncestr
	jsapiTick.Timestamp = timestamp

	// 步骤2. 对string1进行sha1签名，得到signature：
	s := sha1.New()
	io.WriteString(s, string1)

	jsapiTick.Signature = fmt.Sprintf("%x", s.Sum(nil))

	return jsapiTick
}
