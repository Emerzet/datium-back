package httpapi

import (
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Datium")
	w.WriteHeader(http.StatusOK)
	log.Printf("health: %s %s %d", r.Method, r.URL.Path, http.StatusOK)

	w.Write([]byte("ok"))

}
