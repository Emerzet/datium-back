package httpapi

import (
	"net/http"
	"path/filepath"
	"time"
	"strings"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(staticDir string, maxBodyBytes int64) http.Handler {

	r := chi.NewRouter()
	// Global middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * time.Second))
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			req.Body = http.MaxBytesReader(w, req.Body, maxBodyBytes)
			next.ServeHTTP(w, req)

		})

	})

	// Health
	r.Get("/health", Health)
	//Statyka

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, filepath.Join(staticDir, "index.html"))

	})
	//Reszta plików
	fileServer(r, "/", http.Dir(staticDir))
	r.NotFound(func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)

	})

	r.MethodNotAllowed(func(w http.ResponseWriter, req *http.Request) {

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

	})

	return r
}
// Helper fileServer 
func fileServer(r chi.Router, path string, root http.FileSystem) {
if strings.ContainsAny(path, "{}*"){
	panic("fileServer does not permit URL parameters.")

}

if path != "/" && path[len(path)-1] != '/' {
	r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP) 
	path += "/"
	}

r.Get(path+"*", func(w http.ResponseWriter, req *http.Request) {
	rctx := chi.RouteContext(req.Context())
	prefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
	http.StripPrefix(prefix, http.FileServer(root)).ServeHTTP(w, req)
})






}





