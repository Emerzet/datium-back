package httpapi

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(staticDir string) http.Handler {

	r := chi.NewRouter()
	// Global middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * time.Second))

	// Health
	r.Get("/health", Health)
	//Statyka

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, filepath.Join(staticDir, "index.html"))

	})
	//Reszta plików
	fs := http.FileServer(http.Dir(staticDir))
	r.Handle("/*", fs)

	r.NotFound(func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)

	})

	r.MethodNotAllowed(func(w http.ResponseWriter, req *http.Request) {

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

	})

	return r
}
