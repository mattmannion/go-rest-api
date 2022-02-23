package router

import (
	controller_root "mm/pkg/src/controllers/root"
	"mm/pkg/src/middleware"
	"net/http"
)

var mux *http.ServeMux
var Router http.Handler

func init() {
	mux = http.NewServeMux()
	mux.HandleFunc("/", controller_root.RootRouter)

	Router = middleware.MW(mux)
}
