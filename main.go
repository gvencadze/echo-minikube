package main

import (
	"github.com/gvencadze/echo-minikube/internal"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to run: %s", err.Error())
		return
	}
}

func run() error {
	http.HandleFunc("/hello", internal.Hello)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return err
	}

	return nil
}
