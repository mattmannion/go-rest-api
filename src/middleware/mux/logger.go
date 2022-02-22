package mux

import (
	"log"
	"net/http"
)

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("-", r.URL.Path, "-", r.Method)
		h.ServeHTTP(w, r)
	})
}
