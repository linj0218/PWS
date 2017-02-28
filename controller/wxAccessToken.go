package controller

import (
	"PWS/entity"
	"PWS/service"
	"encoding/json"
	"fmt"

	restful "github.com/emicklei/go-restful"
)

// UpdateAccessToken 测试 获取微信access_token
func UpdateAccessToken(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()

	jsoncallback := req.Request.FormValue("jsoncallback")

	resp := entity.ResEntity{}
	err := service.UpdateAccessToken()
	if err != nil {
		resp.Status = 1
		resp.Msg = "失败"
		resp.Data = ""
	} else {
		resp.Status = 0
		resp.Msg = "成功"
		resp.Data = ""
	}

	if jsoncallback != "" {
		responsejson, _ := json.Marshal(resp)
		back := jsoncallback + "(" + string(responsejson) + ")"
		fmt.Fprintln(res.ResponseWriter, back)
	} else {
		res.WriteEntity(resp)
	}
}
