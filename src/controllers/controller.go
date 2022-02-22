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

			case http.MethodOptions:
				options(w, r)

			case http.MethodHead:
				head(w, r)

			default:
				nf(w, r)
			}
		})
}

type data struct {
	Status string `json:"status"`
}

func nf(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(data{Status: "endpoint not available"})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	w.Write(json)
}

func options(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(data{Status: "options"})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}

func head(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }
