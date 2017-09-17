package conf

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

type Route struct {
	Match string
	Methods *[]string
	HandleFunc http.HandlerFunc
}

type Controller struct {
	Name string
	Path string
	Routes *[]Route
}

func RegisterController(controller *Controller) {
	path := controller.Path
	log.Printf("register controller => %s.\n", controller.Name)
	routes := controller.Routes
	for i := 0; i < len(*routes); i++ {
		route := (*routes)[i]
		match := path + route.Match
		methods := route.Methods
		log.Printf("register route match => %s methods => %s.\n", match, *methods)
		Router.HandleFunc(match, route.HandleFunc).Methods(*methods...)
	}
}