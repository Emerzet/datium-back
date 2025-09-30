package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"time"
	"syscall"
	"github.com/Emerzet/datium-back/internal/config"
	"github.com/Emerzet/datium-back/internal/httpserver"
	"github.com/Emerzet/datium-back/internal/infra/httpapi"
	"os"
)

func main() {
	cfg := config.Load()

	log.Println("Serving static from:", cfg.App.StaticDir)
	router := httpapi.NewRouter(cfg.App.StaticDir, cfg.HTTP.MaxBodyBytes)
	srv := httpserver.New(cfg.HTTP, router)

	log.Println("Listening on", cfg.HTTP.Addr)
	
	
	go func(){
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatal("http server error", err)

	}
}()
sigCh := make(chan os.Signal, 1)
signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
<-sigCh

ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
defer cancel()

if err := srv.Shutdown(ctx); err != nil {
	log.Println("graceful shutdown error", err)

}else {
	log.Println("serrver stopped cleanly")
}



}
