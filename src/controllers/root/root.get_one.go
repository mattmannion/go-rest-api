package controller_root

import (
	"encoding/json"
	"log"
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
