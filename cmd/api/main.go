package main

import (
	"net/http"
	"log"
	"github.com/Emerzet/datium-back/internal/config"
	"github.com/Emerzet/datium-back/internal/httpx"
)



func main() {
	cfg := config.Load()
	router := httpx.NewRouter()

	srv := &http.Server{
	Addr: cfg.Addr,
	Handler: router,


	}


	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	log.Fatal(err)

	}


}
