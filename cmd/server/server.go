package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/b3nhard/chitempl/web/components"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Mount("/debug", middleware.Profiler())
	dir, _ := os.Getwd()
	// log.Printf("Working Directory: %s", )
	fs := http.Dir(path.Join(dir, "web/static/"))
	log.Println("FSS:=> ", fs)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		components.Home().Render(context.Background(), w)
	})

	FileServer(r, "/static", fs)

	log.Println("Server starting at https://localhost:4000")
	log.Fatal(http.ListenAndServeTLS(":4000", "server.crt", "server.key", r))
	// log.Fatal(http.ListenAndServe(":4000", r))
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
