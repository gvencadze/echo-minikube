package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Hello endpoint is greeting user by his name passed in URL query parameter
func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("name parameter is required")
		return
	}

	if _, err := fmt.Fprintf(w, "Hello from server in minikube, %s", name); err != nil {
		log.Println(err)
		return
	}
}
