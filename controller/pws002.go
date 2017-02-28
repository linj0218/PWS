package controller

import (
	"encoding/json"
	"fmt"

	"PWS/service"

	restful "github.com/emicklei/go-restful"
)

// PWS002 ...
func PWS002(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()

	jsoncallback := req.Request.FormValue("jsoncallback")

	resp := service.PWS002()

	if jsoncallback != "" {
		responsejson, _ := json.Marshal(resp)
		back := jsoncallback + "(" + string(responsejson) + ")"
		fmt.Fprintln(res.ResponseWriter, back)
	} else {
		res.WriteEntity(resp)
	}
}
