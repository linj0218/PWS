package controller

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"PWS/entity"

	restful "github.com/emicklei/go-restful"
)

// PWS003 ...
func PWS003(req *restful.Request, res *restful.Response) {
	req.Request.ParseForm()
	body, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		fmt.Println("err")
		fmt.Fprintln(res.ResponseWriter, "系统出错!")
	}
	fmt.Println("body: ", string(body))
	reqBody := &entity.WxTextRequest{}
	xml.Unmarshal(body, reqBody)
	fmt.Fprintln(res.ResponseWriter, reqBody)
}
