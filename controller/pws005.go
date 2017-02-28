package controller

import (
	"PWS/entity"
	"PWS/service"
	"encoding/json"
	"fmt"

	restful "github.com/emicklei/go-restful"
)

// PWS005 生成jssdk签名
func PWS005(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()
	url := req.Request.FormValue("url")
	jsoncallback := req.Request.FormValue("jsoncallback")

	ticket := service.PWS005(url)

	resp := entity.ResEntity{}
	resp.Status = 0
	resp.Msg = "成功"
	resp.Data = ticket

	if jsoncallback != "" {
		responsejson, _ := json.Marshal(resp)

		back := jsoncallback + "(" + string(responsejson) + ")"

		fmt.Fprintln(res.ResponseWriter, back)
	} else {
		res.WriteEntity(resp)
	}
}
