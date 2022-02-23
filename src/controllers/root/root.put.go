package root

import (
	"encoding/json"
	"mm/pkg/src/models"
	"net/http"
)

func put(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(models.RootDataPut)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
