package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type metric struct {
	Value int `json:"metric"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/app/clientAuth/metric/{mid}", func(w http.ResponseWriter, r *http.Request) {
		log.Print("clientAuth auth custom called")

		marshal, _ := json.Marshal(metric{Value: 1})
		_, err := w.Write(marshal)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}).Methods("GET")

	caCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile("cert/ca.pem")
	if err != nil {
		log.Print(err)
	}
	caCertPool.AppendCertsFromPEM(pem)
	tlsConf := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	tlsConf.BuildNameToCertificate()
	server := &http.Server{
		Addr:      ":6000",
		TLSConfig: tlsConf,
	}

	if err := server.ListenAndServeTLS("cert/server.pem", "cert/server.key"); err != nil {
		panic(err)
	}
}
