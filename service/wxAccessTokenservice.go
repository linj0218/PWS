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

// AccessTokenResponse 获取access_token正确返回对象
type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

// AccessTokenErrorResponse 获取access_token错误返回对象
type AccessTokenErrorResponse struct {
	Errcode float64
	Errmsg  string
}

// FetchAccessToken 获取access_token
// 拼接get请求 解析返回json结果 返回 access_token和err
func FetchAccessToken(appID, appSecret string) (string, error) {

	requestLine := strings.Join([]string{"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=",
		appID,
		"&secret=",
		appSecret}, "")

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("发送get请求获取 atoken 错误: %s", err)
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if !bytes.Contains(body, []byte("access_token")) {
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		fmt.Printf("发送get请求获取 微信返回 的错误信息 %+v\n", ater)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("%s", ater.Errmsg)
	}

	atr := AccessTokenResponse{}
	err = json.Unmarshal(body, &atr)
	if err != nil {
		log.Printf("发送get请求获取 atoken 返回数据json解析错误: %s", err)
		return "", err
	}
	return atr.AccessToken, nil
}

// UpdateAccessToken 更新access_token
func UpdateAccessToken() error {
	// 使用 appid, appSecret 获取access_token
	log.Printf("开始更新 access_token ...\n")

	appID := util.WXAPPID
	appSecret := util.WXAPPSECRET

	accessToken, _ := FetchAccessToken(appID, appSecret)

	GetConn()
	updateSQL := `update wx_pub_info set accesstoken=? where appid=?`
	_, err := db.Exec(updateSQL, accessToken, appID)
	if err != nil {
		log.Printf("access_token更新失败: %s", err)
		return err
	}
	log.Printf("access_token更新成功!")
	return nil
}

// GetAccessToken 获取access_token
func GetAccessToken() string {
	GetConn()

	appID := util.WXAPPID
	var accessToken string

	selectSQL := `select accesstoken from wx_pub_info where appid=?`
	err := db.Get(&accessToken, selectSQL, appID)
	if err != nil {
		fmt.Println("数据库查询错误: ", err)
	}
	return accessToken
}
