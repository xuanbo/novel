package main

import (
	"log"
	"net/http"
	"github.com/xuanbo/novel/web/conf"
	"github.com/xuanbo/novel/web/controller"
)

func main() {
	conf.UseNotFound()
	conf.StaticResource("/static/ui/dist", "/static/")
	conf.RegisterController(controller.UserController)

	addr := ":9000"
	log.Printf("start server on %s.\n", addr)
	log.Fatal(http.ListenAndServe(addr, conf.Router))
}