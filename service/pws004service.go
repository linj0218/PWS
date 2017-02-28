package service

import (
	"PWS/entity"
	"PWS/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// PWS004 获取用户信息 重定向
func PWS004(code string) string {

	if code == "" {
		return ""
	}
	// 获取access_token
	appID := util.WXAPPID
	appSecret := util.WXAPPSECRET
	requestLine := strings.Join([]string{"https://api.weixin.qq.com/sns/oauth2/access_token",
		"?appid=", appID,
		"&secret=", appSecret,
		"&code=", code,
		"&grant_type=authorization_code"}, "")

	res, err := http.Get(requestLine)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Printf("发送get请求获取 atoken 错误: %s", err)
		return ""
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	atr := entity.AccessToken2UserInfo{}
	_ = json.Unmarshal(body, &atr)

	// 用access_token和openid换取用户信息
	requestLine2 := strings.Join([]string{"https://api.weixin.qq.com/sns/userinfo",
		"?access_token=", atr.AccessToken,
		"&openid=", atr.OpenID,
		"&lang=zh_CN"}, "")

	res2, err2 := http.Get(requestLine2)
	if err2 != nil || res2.StatusCode != http.StatusOK {
		log.Printf("发送get请求获取用户信息错误: %s", err)
		return ""
	}
	defer res2.Body.Close()

	body2, _ := ioutil.ReadAll(res2.Body)
	wxUserInfo := entity.WxUserInfo{}
	_ = json.Unmarshal(body2, &wxUserInfo)

	fmt.Println("userinfo: ", wxUserInfo)

	// 保存用户信息
	openID := wxUserInfo.OpenID

	var wuiid int32
	GetConn()
	selectSQL := `select wui_id from wx_user_info where openid=? and delflg=0`
	_ = db.Get(&wuiid, selectSQL, openID)

	if wuiid == 0 {
		log.Printf("未查询到openid: %s 的数据，数据插入中...", openID)

		insertSQL := `insert into wx_user_info(openid,nickname,sex,language,city,province,country,headimgurl,createtime) values(?,?,?,?,?,?,?,?,?)`

		sqlReturn, _ := db.Prepare(insertSQL)
		defer sqlReturn.Close()

		_, err4 := sqlReturn.Exec(wxUserInfo.OpenID, wxUserInfo.NickName, wxUserInfo.Sex, wxUserInfo.Language, wxUserInfo.City, wxUserInfo.Province, wxUserInfo.Country, wxUserInfo.HeadImgURL, util.GenerateTime())
		if err4 != nil {
			log.Printf("数据库插入失败,err: %s", err4)
			return ""
		}
		log.Printf("用户 %s 信息插入成功", openID)
	}

	return openID
}
