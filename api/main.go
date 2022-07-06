package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"

	"github.com/kimai07/openapi-go-demo/api/generated/openapi"
	"github.com/kimai07/openapi-go-demo/api/server"
	embed "github.com/kimai07/openapi-go-demo/openapi"
)

func addHandlersForOpenAPI(r *chi.Mux)  {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	r.Handle("/api-docs", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(swagger); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))

	fsys, _ := fs.Sub(embed.Content, "swagger-ui")
	r.Handle("/swagger-ui/*", http.StripPrefix("/swagger-ui", http.FileServer(http.FS(fsys))))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	addHandlersForOpenAPI(r)

	s := server.NewServer()
	openapi.HandlerFromMux(s, r)

	if err := http.ListenAndServe(":8080", r); err != nil { // 8080ポートをリッスン
		panic(1)
	}
}
