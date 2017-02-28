package main

import (
	"PWS/controller"
	"fmt"
	"log"
	"net/http"
	"path"

	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

var (
	serviceurl  = "http://127.0.0.1"
	port        = "8888"
	rootdir     = "./views"
	rootHTMLdir = "./views/html"
	rootIMGdir  = "./views/img"
	rootCSSdir  = "./views/css"
	rootJSdir   = "./views/js"
)

// Register ...
func Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/pws").
		Doc("Manage pws").
		Consumes(restful.MIME_JSON, "application/x-www-form-urlencoded", "multipart/form-data", "text/xml", "text/plain").
		Produces(restful.MIME_JSON, "application/json;charset=UTF-8", "text/plain") // you can specify this per route as well
	// 静态文件服务
	ws.Route(ws.GET("/html/{subpath:*}").Filter(encodingFilter).To(dealHTMLPath))
	ws.Route(ws.GET("/img/{subpath:*}").Filter(encodingFilter).To(dealIMGPath))
	ws.Route(ws.GET("/css/{subpath:*}").Filter(encodingFilter).To(dealCSSPath))
	ws.Route(ws.GET("/js/{subpath:*}").Filter(encodingFilter).To(dealJSPath))
	ws.Route(ws.GET("/MP_verify_hTE8AtW0XIZyyH9q.txt").Filter(encodingFilter).To(func(req *restful.Request, resp *restful.Response) {
		actual := path.Join(rootdir, "MP_verify_hTE8AtW0XIZyyH9q.txt")
		http.ServeFile(resp.ResponseWriter, req.Request, actual)
		fmt.Println(resp.ResponseWriter)
	}))

	// 测试
	ws.Route(ws.GET("/test").Filter(webserviceLogging).To(controller.Test))
	// 测试2
	ws.Route(ws.GET("/test2").Filter(webserviceLogging).To(controller.GetUserList))
	// 测试接口pws001
	ws.Route(ws.GET("/pws001").Filter(webserviceLogging).To(controller.PWS001))
	// 测试解析excel
	ws.Route(ws.GET("/pws002").Filter(webserviceLogging).To(controller.PWS002))
	// 测试post请求,测试接口
	ws.Route(ws.GET("/pws003").Filter(webserviceLogging).To(controller.PWS003))
	// 接入微信
	ws.Route(ws.GET("/wx_connect").Filter(webserviceLogging).To(controller.WxConnectController))
	// 微信回复消息
	ws.Route(ws.POST("/wx_connect").Filter(webserviceLogging).To(controller.WxResponseMsg))
	// 微信更新 access_token
	ws.Route(ws.GET("/updateAccessToken").Filter(webserviceLogging).To(controller.UpdateAccessToken))
	// 获取用户信息 重定向
	ws.Route(ws.GET("/pws004").Filter(webserviceLogging).To(controller.PWS004))
	// 获取jsapi权限
	ws.Route(ws.GET("/pws005").Filter(webserviceLogging).To(controller.PWS005))
	// 更新jsapi_ticket
	ws.Route(ws.GET("/pws006").Filter(webserviceLogging).To(controller.PWS006))
	// 设置微信自定义菜单
	ws.Route(ws.GET("/pws007").Filter(webserviceLogging).To(controller.PWS007))
	// 用户信息
	ws.Route(ws.GET("/pws008").Filter(webserviceLogging).To(controller.PWS008))
	// 砍价
	ws.Route(ws.GET("/pws009").Filter(webserviceLogging).To(controller.PWS009))
	// 卡券列表
	ws.Route(ws.GET("/pws010").Filter(webserviceLogging).To(controller.PWS010))
	// 抽奖
	ws.Route(ws.GET("/pws011").Filter(webserviceLogging).To(controller.PWS011))
	// 使用卡券

	container.Add(ws)
}

func encodingFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[encoding-filter] %s,%s\n", req.Request.Method, req.Request.URL)
	// wrap responseWriter into a compressing one
	compress, _ := restful.NewCompressingResponseWriter(resp.ResponseWriter, restful.ENCODING_GZIP)
	resp.ResponseWriter = compress
	defer func() {
		compress.Close()
	}()
	chain.ProcessFilter(req, resp)
}
func dealHTMLPath(req *restful.Request, resp *restful.Response) {
	actual := path.Join(rootHTMLdir, req.PathParameter("subpath"))
	http.ServeFile(resp.ResponseWriter, req.Request, actual)
}
func dealIMGPath(req *restful.Request, resp *restful.Response) {
	actual := path.Join(rootIMGdir, req.PathParameter("subpath"))
	// fmt.Printf("serving %s ... (from %s)\n", actual, req.PathParameter("subpath"))
	http.ServeFile(resp.ResponseWriter, req.Request, actual)
}
func dealCSSPath(req *restful.Request, resp *restful.Response) {
	actual := path.Join(rootCSSdir, req.PathParameter("subpath"))
	http.ServeFile(resp.ResponseWriter, req.Request, actual)
}
func dealJSPath(req *restful.Request, resp *restful.Response) {
	actual := path.Join(rootJSdir, req.PathParameter("subpath"))
	http.ServeFile(resp.ResponseWriter, req.Request, actual)
}

// 服务器日志
func webserviceLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[request info:] %s\n", req.Request.URL)
	chain.ProcessFilter(req, resp)
}

func main() {
	wsContainer := restful.NewContainer()
	Register(wsContainer)

	config := swagger.Config{
		WebServices:     wsContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl:  serviceurl + ":" + port,
		ApiPath:         "/apidocs.json",
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "/Users/ljpc/goworkspace/swagger-ui/dist"}
	swagger.RegisterSwaggerService(config, wsContainer)

	log.Printf("start listening on %s:%s", serviceurl, port)
	server := &http.Server{Addr: ":" + port, Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
