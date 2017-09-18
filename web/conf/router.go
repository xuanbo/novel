package conf

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

// 定义路由
type Route struct {
	Match string
	Methods []string
	HandleFuncName string
	HandleFunc http.HandlerFunc
}

type Routes []Route

// 定义控制器
type Controller struct {
	Name string
	Path string
	Routes *Routes
}

// 注册控制器
func RegisterController(controller *Controller) {
	path := controller.Path
	log.Printf("register controller: %s.\n", controller.Name)
	routes := controller.Routes
	for i := 0; i < len(*routes); i++ {
		route := (*routes)[i]
		match := path + route.Match
		methods := route.Methods
		handleFuncName := route.HandleFuncName
		// 包装打印log
		handleFunc := LoggerHandlerFunc(route.HandleFunc, handleFuncName)
		// 包装跨域
		handleFunc = CorsHandlerFunc(handleFunc)
		log.Printf("register route: %s\t%s\t%s.\n", methods, match, handleFuncName)
		Router.HandleFunc(match, handleFunc).Methods(methods...)
	}
}

func StaticResource(pathPrefix string) {
	Router.PathPrefix(pathPrefix).Handler(http.FileServer(http.Dir(".")))
}