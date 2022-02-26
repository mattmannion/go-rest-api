package mux_mw

import "net/http"

func Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")
		whitelist := []string{"http://localhost:3000", "*"}

		var match string

		switch origin {
		case whitelist[0]:
			match = whitelist[0]
		default:
			match = whitelist[len(whitelist)-1]
		}

		w.Header().Set("Access-Control-Allow-Origin", match)
		w.Header().Set("Access-Control-Allow-Methods", "*")

		h.ServeHTTP(w, r)
	})
}
