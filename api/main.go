package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"

	"github.com/kimai07/openapi-go-demo/api/generated/openapi"
	"github.com/kimai07/openapi-go-demo/api/server"
	embed "github.com/kimai07/openapi-go-demo/openapi"
)

func oapiDocHandler(endpoint string, swagger *openapi3.T) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" && strings.EqualFold(r.URL.Path, endpoint) {
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(swagger)
				if err != nil {
					panic(1)
				}
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func main() {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(oapiDocHandler("/api-docs", swagger))

	fsys, _ := fs.Sub(embed.Content, "swagger-ui")
	r.Handle("/swagger-ui/*", http.StripPrefix("/swagger-ui", http.FileServer(http.FS(fsys))))

	s := server.NewServer()
	openapi.HandlerFromMux(s, r)

	err = http.ListenAndServe(":8080", r) // 8080ポートをリッスン
	if err != nil {
		panic(1)
	}
}
