package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gvencadze/echo-minikube/internal"
	"go.elastic.co/apm/module/apmgorilla/v2"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to run: %s", err.Error())
		return
	}
}

func run() error {
	r := mux.NewRouter()
	apmgorilla.Instrument(r)

	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/hello", internal.Hello)

	err := http.ListenAndServe(":7000", r)
	if err != nil {
		return err
	}

	return nil
}
