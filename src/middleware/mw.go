package middleware

import (
	"mm/pkg/src/middleware/mux"
	"net/http"
)

/* This function is at the top if the file for easy of use */
func MW(h http.Handler) http.Handler {
	mw := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	// add new mw here
	return mux_mw(mw, mux.Cors, mux.Logger)
}

/* mw loop state for mux */
var mux_mw_init bool = false

type mux_mw_chain func(http.Handler) http.Handler

func mux_mw(hf http.Handler, mmc ...mux_mw_chain) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !mux_mw_init {
			for _, mx := range mmc {
				hf = mx(hf)
			}
			mux_mw_init = true
		}

		hf.ServeHTTP(w, r)
	})
}

/* mw loop state for routes */
var route_mw_init bool = false

type route_mw_chain func(http.HandlerFunc) http.HandlerFunc

func Route_mw(hf http.HandlerFunc, rmc ...route_mw_chain) http.HandlerFunc {
	if !route_mw_init {
		for _, mx := range rmc {
			hf = mx(hf)
		}
		route_mw_init = true
	}
	return hf
}
