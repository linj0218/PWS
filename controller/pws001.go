package controller

import (
	"encoding/json"
	"fmt"

	"PWS/entity"

	restful "github.com/emicklei/go-restful"
)

type userInfo0 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PWS001 ...
func PWS001(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()
	test := req.Request.FormValue("test")
	fmt.Println("test: ", test)

	jsoncallback := req.Request.FormValue("jsoncallback")

	userInfo := userInfo0{ID: 1, Name: "小五"}

	resp := entity.ResEntity{}
	resp.Status = 0
	resp.Msg = "成功"
	resp.Data = userInfo

	if jsoncallback != "" {
		responsejson, _ := json.Marshal(resp)

		back := jsoncallback + "(" + string(responsejson) + ")"

		fmt.Fprintln(res.ResponseWriter, back)
	} else {
		res.WriteEntity(resp)
	}
}
