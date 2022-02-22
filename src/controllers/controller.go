package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type controller http.HandlerFunc

type Methods struct {
	Get    controller
	Post   controller
	Put    controller
	Patch  controller
	Delete controller
}

func ControlSwitch(m Methods) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			method := r.Method

			switch method {
			case http.MethodGet:
				if m.Get == nil {
					m.Get = nf
				}
				m.Get(w, r)

			case http.MethodPost:
				if m.Post == nil {
					m.Post = nf
				}
				m.Post(w, r)

			case http.MethodPut:
				if m.Put == nil {
					m.Put = nf
				}
				m.Put(w, r)

			case http.MethodPatch:
				if m.Patch == nil {
					m.Patch = nf
				}
				m.Patch(w, r)

			case http.MethodDelete:
				if m.Delete == nil {
					m.Delete = nf
				}
				m.Delete(w, r)

			default:
				nf(w, r)
			}
		})
}

type not_found struct {
	Status string `json:"status"`
}

func nf(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(not_found{Status: "No endpoint found"})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	w.Write(json)
}
