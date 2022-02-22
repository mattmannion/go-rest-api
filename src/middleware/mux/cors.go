package mux

import "net/http"

func Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")
		wl := []string{"http://localhost:3000", "*"}

		var match string

		switch origin {
		case wl[0]:
			match = wl[0]
		default:
			match = wl[len(wl)-1]
		}

		w.Header().Set("Access-Control-Allow-Origin", match)
		w.Header().Set("Access-Control-Allow-Methods", "*")

		h.ServeHTTP(w, r)
	})
}
