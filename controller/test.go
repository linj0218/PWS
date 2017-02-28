package controller

import (
	"encoding/json"
	"fmt"

	"github.com/emicklei/go-restful"
)

type resData struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
type userInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Test ..
func Test(req *restful.Request, res *restful.Response) {

	res.WriteEntity("test333")
}

// GetUserList 测试接口2
func GetUserList(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()
	jsoncallback := req.Request.FormValue("jsoncallback")
	userInfo := userInfo{ID: 1, Name: "小五"}

	resp := resData{}

	resp.Status = 1
	resp.Data = userInfo

	if jsoncallback != "" {
		responsejson, _ := json.Marshal(resp)

		back := jsoncallback + "(" + string(responsejson) + ")"

		fmt.Fprintln(res.ResponseWriter, back)
	} else {
		res.WriteEntity(resp)
	}
}
