package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
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
			method := strings.ToUpper(r.Method)

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
