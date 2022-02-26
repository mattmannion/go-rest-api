package router

import (
	"mm/pkg/src/controllers/users"
	"mm/pkg/src/middleware"
	"mm/pkg/src/middleware/mux_mw"
	"net/http"
)

var mux *http.ServeMux
var Router http.Handler

func init() {
	mux = http.NewServeMux()
	mux.HandleFunc("/users", users.Router)

	Router = middleware.Mux_mw(mux, mux_mw.Cors, mux_mw.Logger)
}
