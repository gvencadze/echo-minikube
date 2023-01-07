package main

import (
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

	r.HandleFunc("/hello", internal.Hello)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return err
	}

	return nil
}
