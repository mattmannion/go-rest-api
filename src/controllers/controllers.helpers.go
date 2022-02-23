package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type req_body struct {
	ID int
}

type fn_get_one func(w http.ResponseWriter, r *http.Request, id int)

func GetByID(get_one fn_get_one, get_all http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
	var rb req_body
	json.NewDecoder(r.Body).Decode(&rb)

	if rb.ID > 0 {
		get_one(w, r, rb.ID)
		return
	}

	get_all(w, r)
}

func ResultNotFound(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(map[string]string{"err": "Result Not Found"})
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	w.Write(json)
}
