package main

import (
	"github.com/xuanbo/novel/web/conf"
	"github.com/xuanbo/novel/web/controller"
	"net/http"
)

func main() {
	conf.RegisterController(controller.UserController)
	http.ListenAndServe(":9000", conf.Router)
}