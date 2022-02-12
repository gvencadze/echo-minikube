package internal

import (
	"fmt"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if _, err := fmt.Fprintf(w, "Hello from server in minikube, %s", name); err != nil {
		log.Println(err)
		return
	}
}
