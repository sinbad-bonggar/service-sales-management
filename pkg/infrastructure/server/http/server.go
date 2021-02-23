package httpserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	Router chi.Router
}

// New ...
func New() *Server {
	s := &Server{}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	s.Router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
