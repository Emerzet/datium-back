package main

import (
	"github.com/Emerzet/datium-back/internal/config"
	"github.com/Emerzet/datium-back/internal/infra/httpapi"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	log.Println("Serving static from:", cfg.App.StaticDir)
	router := httpapi.NewRouter(cfg.App.StaticDir)

	srv := &http.Server{
		Addr:              cfg.HTTP.Addr,
		Handler:           router,
		ReadHeaderTimeout: cfg.HTTP.ReadTimeout,
		WriteTimeout:      cfg.HTTP.WriteTimeout,
		IdleTimeout:       cfg.HTTP.IdleTimeout,
	}

	log.Println("Listening on", cfg.HTTP.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)

	}

}
