package apiserver

import (
	"github.com/gorilla/mux"
	"net/http"
	"payment-service/internal/handlers/handlers"
	"payment-service/internal/middleware"
)

type Server struct {
	router     *mux.Router
	handler    *handlers.Handlers
	middleware *middleware.Auth
}

func newServer(router *mux.Router, handler *handlers.Handlers, middleware *middleware.Auth) *Server {
	s := &Server{
		router:     router,
		handler:    handler,
		middleware: middleware,
	}

	s.configureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {

	s.router.HandleFunc("/create-payment", s.handler.Payment().Create()).Methods("POST")
	s.router.HandleFunc("/change-payment-status", s.handler.Payment().ChangeStatus()).Methods("PATCH")
	s.router.HandleFunc("/get-payment-status/{id}", s.handler.Payment().GetStatus()).Methods("GET")
	s.router.HandleFunc("/get-payments/{path}", s.handler.Payment().GetByEmailOrID()).Methods("GET")
	s.router.HandleFunc("/cancel-payment/{id}", s.handler.Payment().Cancel()).Methods("PATCH")
	s.router.Use(s.middleware.BasicAuth)
}
