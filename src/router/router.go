package router

import (
	"mm/pkg/src/controllers/root"
	"mm/pkg/src/controllers/users"
	"mm/pkg/src/middleware"
	"net/http"
)

var mux *http.ServeMux
var Router http.Handler

func init() {
	mux = http.NewServeMux()
	mux.HandleFunc("/users", users.Router)
	mux.HandleFunc("/", root.Router)

	Router = middleware.MW(mux)
}
