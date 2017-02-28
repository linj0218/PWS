package service

import (
	"PWS/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type jsapiTicketResponse struct {
	Errcode   int32  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int32  `json:"expires_in"`
}
type jsapiTicketErrResponse struct {
	Errcode int32  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// PWS006 更新jsapi_ticket
func PWS006() error {

	GetConn()

	// 获取access_token
	appid := util.WXAPPID
	var accesstoken string

	selectSQL := `select accesstoken from wx_pub_info where appid=? and delflg=0`
	err := db.Get(&accesstoken, selectSQL, appid)
	if err != nil {
		log.Printf("数据库查询错误: %s", err)
		return err
	}

	// 发送get请求获取jsapi_ticket
	requestLine := strings.Join([]string{"https://api.weixin.qq.com/cgi-bin/ticket/getticket",
		"?access_token=", accesstoken,
		"&type=jsapi"}, "")

	resp, err2 := http.Get(requestLine)
	if err2 != nil || resp.StatusCode != http.StatusOK {
		log.Printf("发送get请求获取 jsapi_ticket 错误: %s", err2)
		return err2
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if !bytes.Contains(body, []byte("ticket")) {
		ater := jsapiTicketErrResponse{}
		err = json.Unmarshal(body, &ater)
		log.Printf("发送get请求获取 微信返回 的错误信息 %+v\n", ater)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", ater.Errmsg)
	}
	atr := jsapiTicketResponse{}
	err = json.Unmarshal(body, &atr)
	if err != nil {
		log.Printf("发送get请求获取 jsapi_ticket 返回数据json解析错误: %s", err)
		return err
	}

	log.Printf("开始更新 jsapi_ticket ... \n")

	GetConn()
	updateSQL := `update wx_pub_info set jsapiticket=? where appid=?`
	_, err3 := db.Exec(updateSQL, atr.Ticket, appid)
	if err3 != nil {
		log.Printf("jsapi_ticket更新失败: %s", err3)
		return err3
	}
	log.Printf("jsapi_ticket更新成功!")

	return nil
}

// GetJsAPITicket 从数据库读取jsapi_ticket
func GetJsAPITicket() string {
	GetConn()
	appid := util.WXAPPID
	var jsapiticket string
	selectSQL := `select jsapiticket from wx_pub_info where appid=?`
	err := db.Get(&jsapiticket, selectSQL, appid)
	if err != nil {
		fmt.Println("数据库查询错误: ", err)
		return ""
	}
	return jsapiticket
}
