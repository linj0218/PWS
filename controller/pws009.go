package controller

import (
	"PWS/entity"
	"PWS/service"
	"encoding/json"
	"fmt"
	"log"

	restful "github.com/emicklei/go-restful"
)

// PWS009 用户信息
func PWS009(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()

	openid := req.Request.FormValue("openid")
	topenid := req.Request.FormValue("topenid") // 砍价目标openid
	jsoncallback := req.Request.FormValue("jsoncallback")

	info, err := service.PWS009(openid, topenid)

	resp := entity.ResEntity{}
	resp.Status = 0
	resp.Msg = "成功"
	resp.Data = info

	if err != nil {
		log.Printf("砍价失败: %s", err)
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
