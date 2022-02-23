package controller_root

import (
	"encoding/json"
	"log"
	model_root "mm/pkg/src/models/root"
	"net/http"
)

type get_one struct {
	Status string `json:"status"`
	ID     int    `json:"id"`
}

func RootGetOne(w http.ResponseWriter, r *http.Request, id int) {
	json, err := json.Marshal(get_one{Status: "good", ID: id})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}

func RootGetAll(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(model_root.RootDataGetALL)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}

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
