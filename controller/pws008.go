package controller

import (
	"PWS/entity"
	"PWS/service"
	"encoding/json"
	"fmt"

	restful "github.com/emicklei/go-restful"
)

// PWS008 砍价 我的信息
func PWS008(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()
	openid := req.Request.FormValue("openid")
	jsoncallback := req.Request.FormValue("jsoncallback")

	data, err := service.PWS008(openid)

	resp := entity.ResEntity{}
	resp.Status = 0
	resp.Msg = "成功"
	resp.Data = data

	if err != nil {
		resp.Status = 1
		resp.Msg = "失败"
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
