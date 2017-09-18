package main

import (
	"log"
	"net/http"
	"github.com/xuanbo/novel/web/conf"
	"github.com/xuanbo/novel/web/controller"
)

func main() {
	conf.UseNotFound()
	conf.StaticResource("/ui/static")
	conf.RegisterController(controller.UserController)

	addr := ":9000"
	log.Printf("start server on %s.\n", addr)
	http.ListenAndServe(addr, conf.Router)
}