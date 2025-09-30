package httpserver

import (
	"context"
	"net/http"

	"github.com/Emerzet/datium-back/internal/config"
)

type Server struct {
	http *http.Server



}
func New(cfg config.HTTPConfig, handler http.Handler) *Server {
	return &Server{
	http: &http.Server{
		Addr: cfg.Addr,
		Handler: handler,
		ReadTimeout: cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout: cfg.IdleTimeout,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
	},

	}

}


func (s *Server) Start() error {
	return s.http.ListenAndServe()

}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
