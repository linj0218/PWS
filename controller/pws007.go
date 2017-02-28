package controller

import (
	"PWS/entity"
	"PWS/service"
	"encoding/json"
	"fmt"

	restful "github.com/emicklei/go-restful"
)

// PWS007 设置微信自定义菜单
func PWS007(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()
	jsoncallback := req.Request.FormValue("jsoncallback")

	err := service.PWS007()

	resp := entity.ResEntity{}
	resp.Status = 0
	resp.Msg = "成功"
	resp.Data = ""

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
