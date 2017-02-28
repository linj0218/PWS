package controller

import (
	"PWS/entity"
	"PWS/service"
	"encoding/json"
	"fmt"

	restful "github.com/emicklei/go-restful"
)

// PWS010 卡券列表
func PWS010(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()

	openid := req.Request.FormValue("openid")
	jsoncallback := req.Request.FormValue("jsoncallback")

	info, err := service.PWS010(openid)

	resp := entity.ResEntity{}
	resp.Status = 0
	resp.Msg = "成功"
	resp.Data = info

	if err != nil {
		resp.Status = 1
		resp.Msg = "失败"
	}

	if jsoncallback != "" {
		responsejson, _ := json.Marshal(resp)

		back := jsoncallback + "(" + string(responsejson) + ")"

		fmt.Fprintln(res.ResponseWriter, back)
	} else {
		res.WriteEntity(resp)
	}
}
