package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/Emerzet/datium-back/internal/config"
	"github.com/Emerzet/datium-back/internal/httpx"
)



func main() {
	relPublicDir := "../datium-front/public"
	absPublicDir, err := filepath.Abs(relPublicDir)
	if err != nil {
		log.Fatal(err)
	}										
	fs := http.FileServer(http.Dir(absPublicDir))
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
