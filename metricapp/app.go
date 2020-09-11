package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

var apiKey = "SampleApiKey"
var userName = "admin"
var password = "admin"

var customHeaderKey = "X-HEADER-CUSTOM"
var customQueryParamKey = "custom_param"

var stdHeaderKey = "X-API-KEY"
var stdQueryParamKey = "api_key"

type metric struct {
	Value int `json:"metric"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/app/noAuth/metric/{mid}", func(w http.ResponseWriter, r *http.Request) {
		log.Print("no auth called")
		vars := mux.Vars(r)
		metricID := vars["mid"]
		log.Print(metricID)

		marshal, _ := json.Marshal(metric{Value: 1})
		_, err := w.Write(marshal)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}).Methods("GET")

	router.HandleFunc("/app/apiAuth/headers/d/metric/{mid}", func(w http.ResponseWriter, r *http.Request) {
		log.Print("headers default called")
		appKey := r.Header.Get(stdHeaderKey)
		if appKey != apiKey {
			log.Print(r.Header)
			log.Print("key not found default header")
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		marshal, _ := json.Marshal(metric{Value: 1})
		_, err := w.Write(marshal)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}).Methods("GET")

	router.HandleFunc("/app/apiAuth/query/d/metric/{mid}", func(w http.ResponseWriter, r *http.Request) {
		log.Print("query default called")
		appKey := r.URL.Query().Get(stdQueryParamKey)
		if appKey != apiKey {
			log.Print("key not found default query")
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		marshal, _ := json.Marshal(metric{Value: 1})
		_, err := w.Write(marshal)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}).Methods("GET")

	router.HandleFunc("/app/apiAuth/headers/c/metric/{mid}", func(w http.ResponseWriter, r *http.Request) {
		log.Print("header custom called")
		appKey := r.Header.Get(customHeaderKey)
		if appKey != apiKey {
			log.Print("key not found custom header")
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		marshal, _ := json.Marshal(metric{Value: 1})
		_, err := w.Write(marshal)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}).Methods("GET")

	router.HandleFunc("/app/apiAuth/query/c/metric/{mid}", func(w http.ResponseWriter, r *http.Request) {
		log.Print("query custom called")
		appKey := r.URL.Query().Get(customQueryParamKey)
		if appKey != apiKey {
			log.Print("key not found custom query")
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		marshal, _ := json.Marshal(metric{Value: 1})
		_, err := w.Write(marshal)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}).Methods("GET")

	router.HandleFunc("/app/basicAuth/metric/{mid}", func(w http.ResponseWriter, r *http.Request) {
		log.Print("basic auth custom called")
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			log.Print("base auth fail")
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		m := metric{
			Value: 1,
		}

		marshal, _ := json.Marshal(m)
		_, err := w.Write(marshal)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}).Methods("GET")

	if err := http.ListenAndServe(":5000", router); err != nil {
		panic("error starting the listener")
	}
}

func validate(u, p string) bool {
	if u == userName && p == password {
		return true
	}
	return false
}
