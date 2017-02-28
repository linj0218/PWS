package controller

import (
	"net/http"

	"PWS/service"

	restful "github.com/emicklei/go-restful"
)

// PWS004 获取用户信息 重定向
func PWS004(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()
	code := req.Request.FormValue("code")
	state := req.Request.FormValue("state")

	openid := service.PWS004(code)

	var resURL string
	if state == "bargin" {
		resURL = "/pws/html/bargin_index.html?openid=" + openid
	} else {
		resURL = "/pws/html/b_help.html?openid=" + openid + "&topenid=" + state
	}
	http.Redirect(res.ResponseWriter, req.Request, resURL, http.StatusFound)
}
