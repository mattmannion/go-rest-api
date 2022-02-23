package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
)

////////////////////////////////////////
////////// Control Switch //////////////
////////////////////////////////////////

type controller http.HandlerFunc

type Methods struct {
	Get    controller
	Post   controller
	Put    controller
	Patch  controller
	Delete controller
}

func check_controller(c controller, w http.ResponseWriter, r *http.Request) {
	if c == nil {
		c = nf
	}
	c(w, r)
}

func ControlSwitch(m Methods) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			method := strings.ToUpper(r.Method)

			switch method {
			case http.MethodGet:
				check_controller(m.Get, w, r)

			case http.MethodPost:
				check_controller(m.Post, w, r)

			case http.MethodPut:
				check_controller(m.Put, w, r)

			case http.MethodPatch:
				check_controller(m.Patch, w, r)

			case http.MethodDelete:
				check_controller(m.Delete, w, r)

			case http.MethodOptions, http.MethodHead:
				other(w, r)

			default:
				nf(w, r)
			}
		})
}

type data struct {
	Status string `json:"status"`
}

func nf(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(data{Status: "endpoint not available"})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	w.Write(json)
}

func other(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }

///////////////////////////////////////
//////////// Get By ID ////////////////
///////////////////////////////////////

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
