package controller

import (
	"PWS/entity"
	"PWS/service"
	"encoding/json"
	"fmt"

	restful "github.com/emicklei/go-restful"
)

// PWS011 抽奖
func PWS011(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()

	openid := req.Request.FormValue("openid")
	topenid := req.Request.FormValue("topenid")
	jsoncallback := req.Request.FormValue("jsoncallback")

	info, err := service.PWS011(openid, topenid)

	resp := entity.ResEntity{}
	resp.Status = 0
	resp.Msg = "成功"
	resp.Data = info

	if err != nil {
		resp.Status = 1
		resp.Msg = "失败"
		resp.Data = "系统错误！"
	}

	if jsoncallback != "" {
		responsejson, _ := json.Marshal(resp)

		back := jsoncallback + "(" + string(responsejson) + ")"

		fmt.Fprintln(res.ResponseWriter, back)
	} else {
		res.WriteEntity(resp)
	}
}
