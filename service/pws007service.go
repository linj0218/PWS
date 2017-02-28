package service

import (
	"PWS/util"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// PWS007 设置微信自定义菜单
func PWS007() error {
	// 发送post请求，修改微信菜单
	resp, err := http.Post("https://api.weixin.qq.com/cgi-bin/menu/create?access_token="+GetAccessToken(),
		"application/x-www-form-urlencoded",
		strings.NewReader(``+
			`{"button":[`+
			`	{"name":"值日表","type":"view","url":"https://linj0218.github.io/pages/html/rili.html"},`+
			`	{"name":"砍价活动","type":"view","url":"https://open.weixin.qq.com/connect/oauth2/authorize?appid=`+util.WXAPPID+`&redirect_uri=http%3a%2f%2f`+util.SERVICEADDRESS+`%2fpws%2fpws004&response_type=code&scope=snsapi_userinfo&state=bargin#wechat_redirect"},`+
			`	{"name":"其他","sub_button":[`+
			`		{"name":"jssdk示例","type":"view","url":"http://demo.open.weixin.qq.com/jssdk"},`+
			`		{"name":"微信ui库","type":"view","url":"https://weui.io/"}`+
			`	]}`+
			`]}`))

	if err != nil {
		log.Printf("post请求错误: %s", err)
		return err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	log.Printf("微信自定义菜单修改成功! %s", string(body))
	return nil
}
