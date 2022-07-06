package server

import (
	"encoding/json"
	"net/http"

	"github.com/kimai07/openapi-go-demo/api/generated/openapi"
)

type server struct{}

var _ openapi.ServerInterface = (*server)(nil)

func NewServer() *server {
	return &server{}
}

func (s *server) GetHello(w http.ResponseWriter, r *http.Request, params openapi.GetHelloParams) {
	res := s.newHelloResponse(*params.Name)
	s.writeResponse(w, res)
}

func (s *server) newHelloResponse(name string) *openapi.HelloResponse {
	val := "world"
	if name != "" {
		val = name
	}
	val = "hello " + val

	return &openapi.HelloResponse{
		Result: &openapi.Hello{Message: &val},
	}
}

func (s *server) writeResponse(w http.ResponseWriter, res *openapi.HelloResponse) {
	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if _, err = w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
