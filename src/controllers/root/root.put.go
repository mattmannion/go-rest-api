package controller_root

import (
	"encoding/json"
	model_root "mm/pkg/src/models/root"
	"net/http"
)

func RootPut(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(model_root.RootDataPut)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
