package apiserver

import (
	"github.com/gorilla/mux"
	"net/http"
	"payment-service/internal/handlers/handlers"
)

type Server struct {
	router   *mux.Router
	handlers *handlers.Handlers
}

func newServer(router *mux.Router, handlers *handlers.Handlers) *Server {
	s := &Server{
		router:   router,
		handlers: handlers,
	}

	s.configureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/create-transactions", s.handlers.Transact().HandCreatePayment()).Methods("POST")
}
