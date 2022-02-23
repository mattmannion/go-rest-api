package controller_root

import (
	"encoding/json"
	"log"
	model_root "mm/pkg/src/models/root"
	"net/http"
)

func RootPut(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(model_root.RootDataPut)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
