package service

import (
	"PWS/entity"
	"encoding/xml"
	"log"
	"strings"
	"time"
)

// DealMsg 处理微信消息
func DealMsg(body []byte) string {
	var r string
	reqBody := &entity.WxTextRequest{}
	xml.Unmarshal(body, reqBody)

	if reqBody != nil {
		log.Printf("[Wechat Service:] request msg from user [%s] to user [%s], msgType is [%s], request body is \n%s\n",
			reqBody.FromUserName,
			reqBody.ToUserName,
			reqBody.MsgType,
			string(body))

		// 创建回复消息
		if strings.EqualFold(reqBody.MsgType, "text") { // 文本消息
			resBody := makeResponseMsg(reqBody.ToUserName, reqBody.FromUserName, reqBody.Content)
			r = string(resBody)
		} else if strings.EqualFold(reqBody.MsgType, "event") { // 关注、取消关注
			resBody := makeResponseMsg(reqBody.ToUserName, reqBody.FromUserName, reqBody.Event)
			r = string(resBody)
		} else {
			resBody := makeOtherResponseMsg(reqBody.ToUserName, reqBody.FromUserName)
			r = string(resBody)
		}
	}

	return r
}

// makeResponseMsg 创建回复消息
func makeResponseMsg(fromUserName, toUserName, content string) string {
	var str string

	resBody := &entity.WxTextResponse{}
	resBody.FromUserName = entity.Value2CDATA(fromUserName)
	resBody.ToUserName = entity.Value2CDATA(toUserName)
	resBody.MsgType = entity.Value2CDATA("text")

	if strings.EqualFold(content, "日历") {
		resBody.Content = entity.Value2CDATA("https://linj0218.github.io/pages/html/rili.html")

	} else if strings.EqualFold(content, "杰哥") {
		resBody.Content = entity.Value2CDATA("欧桑娇")

	} else if strings.EqualFold(content, "欧桑娇") {
		resBody.Content = entity.Value2CDATA("嘿嘿")

	} else if strings.EqualFold(content, "subscribe") {
		resBody.Content = entity.Value2CDATA("感谢你的关注，该帐号用于测试")

	} else if strings.EqualFold(content, "unsubscribe") {
		resBody.Content = entity.Value2CDATA("感谢你的陪伴，有缘再见")
		log.Println("用户：", toUserName, "已取消关注")

	} else {
		resBody.Content = entity.Value2CDATA("你好啊")

	}

	resBody.CreateTime = time.Duration(time.Now().Unix())
	r, err := xml.Marshal(resBody)

	if err != nil {
		str = "系统异常！"
	} else {
		str = string(r)
	}

	return str
}

// makeOtherResponseMsg 创建未处理消息类型的回复消息
func makeOtherResponseMsg(fromUserName, toUserName string) string {
	var str string

	resBody := &entity.WxTextResponse{}
	resBody.FromUserName = entity.Value2CDATA(fromUserName)
	resBody.ToUserName = entity.Value2CDATA(toUserName)
	resBody.MsgType = entity.Value2CDATA("text")
	resBody.Content = entity.Value2CDATA("暂未对你的请求类型做处理哦")
	resBody.CreateTime = time.Duration(time.Now().Unix())
	r, err := xml.Marshal(resBody)

	if err != nil {
		str = "系统异常！"
	} else {
		str = string(r)
	}

	return str
}
