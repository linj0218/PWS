package entity

// BarginCard 卡券信息
type BarginCard struct {
	Type     int    `json:"type"` // 1砍价,2抽奖
	Name     string `json:"name"`
	CDKey    string `json:"cdkey"`
	IsUsed   int    `json:"isused"`
	UsedTime string `json:"usedtime"`
}

// BarginUserInfo 用户信息
type BarginUserInfo struct {
	NickName    string `json:"nickname"`
	HeadImgURL  string `json:"headimgurl"`
	TargetTime  string `json:"targettime"`
	BbarginTime int    `json:"bbargintime"` // 砍价次数
}

// BarginSQL 砍价记录
type BarginSQL struct {
	NickName   string `json:"nickname"`
	HeadImgURL string `json:"headimgurl"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	CDKey      string `json:"cdkey"`
	Usedtime   string `json:"usedtime"`
	IsUsed     string `json:"isused"`
}
