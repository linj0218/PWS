package controller

import (
	"PWS/service"
	"fmt"
	"io/ioutil"

	restful "github.com/emicklei/go-restful"
)

// WxResponseMsg 微信消息
func WxResponseMsg(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()
	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		fmt.Println("err: ", err)
	}

	resBody := service.DealMsg(body)

	fmt.Fprintln(res.ResponseWriter, resBody)
}
