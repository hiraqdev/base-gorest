package app

import (
	"net/http"

	"github.com/hiraqdev/base-gorest/app/modules/ping"
)

// Routers used to register all available routes
var Routers route

type httpController struct {
	Method  string
	Handler httpHandler
}

type httpHandler func(h http.ResponseWriter, r *http.Request)
type route map[string]httpController

func init() {
	// you should place all of your available routes here
	Routers = make(map[string]httpController)
	Routers["/ping"] = httpController{"GET", ping.Handler}
}
