package controller_root

import (
	"encoding/json"
	"net/http"
)

type req_body struct {
	ID int
}

func RootGet(w http.ResponseWriter, r *http.Request) {
	var rb req_body
	json.NewDecoder(r.Body).Decode(&rb)

	if rb.ID > 0 {
		RootGetOne(w, r, rb.ID)
		return
	}

	RootGetAll(w, r)
}
