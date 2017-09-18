package conf

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/xuanbo/novel/core/model"
	"github.com/xuanbo/novel/web/util"
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
	for _, route := range *routes {
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

type NotFoundHandler struct {
}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(util.ToJsonByte(model.NotFound()))
}

func StaticResource(pathPrefix string) {
	Router.PathPrefix(pathPrefix).Handler(http.FileServer(http.Dir(".")))
}

func UseNotFound() {
	Router.NotFoundHandler = NotFoundHandler{}
}