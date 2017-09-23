package main

import (
	"log"
	"net/http"
	"github.com/xuanbo/novel/web/conf"
	"github.com/xuanbo/novel/web/controller"
)

func main() {
	conf.UseNotFound()
	conf.StaticResource("/static/")
	conf.RegisterController(controller.UserController)

	addr := ":9000"
	log.Printf("listen on %s\n", "http://localhost:9000/static/")
	log.Fatal(http.ListenAndServe(addr, conf.Router))
}