package root

import (
	"encoding/json"
	"log"
	"mm/pkg/src/models"
	"net/http"
)

func get_all(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(models.RootDataGetALL)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
