package mux

import (
	"log"
	"net/http"
	"strings"
)

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("-", r.URL.Path, "-", strings.ToUpper(r.Method))
		h.ServeHTTP(w, r)
	})
}
