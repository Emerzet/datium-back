package httpx

import (
	"net/http"
	"path/filepath"
	"github.com/Emerzet/datium-back/internal/handlers"
)





func NewRouter(staticDir string) http.Handler {
	
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", handlers.Health)
	//Statyka

	
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
		return 
	}


	fs.ServeHTTP(w, r)
	}))


	return mux

}
