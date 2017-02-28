package entity

import "encoding/xml"
import "time"

// WxTextRequest 微信文本请求entity
type WxTextRequest struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string // 小视频为shortvideo,视频为video,语音为voice,图片为image,文本为text,地理位置为location,链接为link
	// 普通消息
	Content      string // 文本消息内容
	MsgID        int
	PicURL       string // 图片链接
	MediaID      string // 图片消息媒体id，语音消息媒体id，视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	Format       string // 语音格式，如amr，speex等
	Recognition  string // 语音识别结果，UTF8编码
	ThumbMediaID string // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据
	LocationX    string // 地理位置维度
	LocationY    string // 地理位置经度
	Scale        string // 地图缩放大小
	Label        string // 地理位置信息
	Title        string // 消息标题
	Description  string // 消息描述
	URL          string // 消息链接
	// 关注
	Event string // subscribe(订阅)、unsubscribe(取消订阅)
}

// WxTextResponse 微信文本回复entity
type WxTextResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
	MediaID      CDATAText
}

// CDATAText 返回xml格式
type CDATAText struct {
	Text string `xml:",innerxml"`
}

// AccessTokenEntity appid和access_token实体
type AccessTokenEntity struct {
	AppID       string
	AccessToken string
}

// Value2CDATA xml转成微信xml格式
func Value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

// AccessToken2UserInfo 获取用户信息前获取的access_token
type AccessToken2UserInfo struct {
	AccessToken  string  `json:"access_token"`
	ExpiresIn    float64 `json:"expires_in"`
	RefreshToken string  `json:"refresh_token"`
	OpenID       string  `json:"openid"`
	Scope        string  `json:"scope"`
}

// WxUserInfo 获取微信用户信息
type WxUserInfo struct {
	OpenID     string `json:"openid"`
	NickName   string `json:"nickname"`
	Sex        int32  `json:"sex"`
	Language   string `json:"language"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	HeadImgURL string `json:"headimgurl"`
}

// JsapiTicket jsapi接口返回
type JsapiTicket struct {
	AppID     string `json:"appId"`
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}
